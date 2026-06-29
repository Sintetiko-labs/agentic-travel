package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
	tkhotel "github.com/fbelchi/travelkit/hotel"
	"github.com/fbelchi/travelkit/parse"
)

// Search queries NH Hotel Group search API, falling back to directory HTML.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	if res, err := c.searchAPI(query, page, pageSize); err == nil && len(res.Hotels) > 0 {
		return res, nil
	} else if err != nil && !shouldFallbackNH(err) {
		return nil, err
	}
	return c.searchDirectory(query, page, pageSize)
}

func shouldFallbackNH(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "akamai blocked") ||
		strings.Contains(msg, "http 403") ||
		strings.Contains(msg, "http 404")
}

func (c *Client) searchAPI(query string, page, pageSize int) (*HotelSearchResult, error) {
	path := fmt.Sprintf("/nh/es/api/v1/hotels/search?query=%s&locale=es&page=%d&size=%d",
		url.QueryEscape(query), page, pageSize)
	var resp nhSearchResponse
	if err := c.GetJSON(path, &resp); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("nh"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	resp.normalize()
	return resp.toResult(query, page, pageSize, c.Brand, c.BaseURL), nil
}

func (c *Client) searchDirectory(query string, page, pageSize int) (*HotelSearchResult, error) {
	parseFn := func(html, base string) []parse.HotelLD {
		return parse.HotelsFromNHDirectory(html, base, query)
	}
	rows, err := tkhotel.SpanishHTMLSearch(c.FetchHTML, c.BaseURL, tkhotel.NHDirectoryPaths(query), parseFn, query)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("nh"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	filtered := tkhotel.FilterSpanishQuery(rows, query)
	if len(filtered) == 0 {
		return nil, fmt.Errorf("search %q: no hotels in directory — %s", query, akamai.NeedsSessionHint("nh"))
	}
	brand := c.Brand
	if brand == "" {
		brand = "NH Hotels"
	}
	return tkhotel.LDToResult(filtered, query, page, pageSize, brand, c.BaseURL, "directory"), nil
}

type nhHotelRow struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	City     string  `json:"city"`
	Country  string  `json:"country"`
	Stars    float64 `json:"stars"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	Slug     string  `json:"slug"`
	Image    string  `json:"image"`
}

type nhSearchResponse struct {
	Total  int          `json:"total"`
	Data   []nhHotelRow `json:"data"`
	Hotels []nhHotelRow `json:"hotels"`
}

func (r *nhSearchResponse) normalize() {
	if len(r.Data) == 0 && len(r.Hotels) > 0 {
		r.Data = r.Hotels
	}
	if r.Total == 0 {
		r.Total = len(r.Data)
	}
}

func (r *nhSearchResponse) toResult(query string, page, pageSize int, brand, base string) *HotelSearchResult {
	hits := make([]HotelHit, 0, len(r.Data))
	for _, h := range r.Data {
		price := ""
		if h.Price > 0 {
			price = fmt.Sprintf("%.2f", h.Price)
		}
		b := h.Brand
		if brand != "" {
			b = brand
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: b, City: h.City, Country: h.Country,
			Stars: h.Stars, Price: price, Currency: h.Currency,
			HotelURL: tkbase.Absolutize(base, "/es/hotel/"+h.Slug),
			ImageURL: tkbase.Absolutize(base, h.Image),
		})
	}
	return &HotelSearchResult{
		Query: query, Total: r.Total, Page: page, PageSize: pageSize,
		HasNext: r.Total > page*pageSize, Hotels: hits, Brand: brand, Source: "api",
	}
}
