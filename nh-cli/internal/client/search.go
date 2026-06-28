package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search queries NH Hotel Group search API.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
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
