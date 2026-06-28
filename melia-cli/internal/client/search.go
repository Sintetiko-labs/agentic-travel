package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
)

const (
	meliaSearchBFF  = "/services/search/hotels/v2/search"
	meliaHotelsPath = "/es/hoteles"
)

// Search queries Meliá hotel search BFF, falling back to the hotel directory page.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	if res, err := c.searchBFF(query, page, pageSize); err == nil {
		return res, nil
	} else if !shouldFallbackSearch(err) {
		return nil, err
	}
	return c.searchDirectory(query, page, pageSize)
}

func shouldFallbackSearch(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "http 404") || strings.Contains(msg, "decode json")
}

func (c *Client) searchBFF(query string, page, pageSize int) (*HotelSearchResult, error) {
	payload := map[string]any{
		"text":     query,
		"language": "es",
		"market":   "ES",
		"page":     page,
		"size":     pageSize,
	}
	body, status, err := c.postMelia(meliaSearchBFF, payload)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		if akamai.IsDenied(status, string(body)) || akamai.IsAppNotFoundWithoutSession(status, string(body)) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("melia"))
		}
		return nil, fmt.Errorf("search %q: HTTP %d: %s", query, status, tkbase.Truncate(string(body), 200))
	}
	resp, err := decodeMeliaSearch(body)
	if err != nil {
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	return resp.toResult(query, page, pageSize, c.Brand, c.BaseURL), nil
}

func (c *Client) searchDirectory(query string, page, pageSize int) (*HotelSearchResult, error) {
	html, err := c.FetchHTML(c.BaseURL + meliaHotelsPath)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("melia"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	rows := parse.HotelsFromMeliaDirectory(html, c.BaseURL, query)
	total := len(rows)
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}
	hits := make([]HotelHit, 0, end-start)
	for _, h := range rows[start:end] {
		b := c.Brand
		if b == "" {
			b = "Meliá"
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: b, City: query,
			Stars: h.Stars, HotelURL: h.URL, ImageURL: h.ImageURL,
		})
	}
	return &HotelSearchResult{
		Query: query, Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Hotels: hits, Brand: c.Brand, Source: "directory",
	}, nil
}

func (c *Client) postMelia(path string, payload any) ([]byte, int, error) {
	c.Throttle()
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}
	req, err := httpNewPost(c.BaseURL+path, b)
	if err != nil {
		return nil, 0, err
	}
	c.SetAPIHeaders(req)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", c.BaseURL)
	req.Header.Set("referer", c.BaseURL+"/es/")
	req.Header.Set("melia-language", "es")
	req.Header.Set("melia-market", "ES")
	req.Header.Set("channel", "web")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := ioReadAll(resp.Body)
	return body, resp.StatusCode, nil
}

func (c *Client) PostRaw(url string, payload any) ([]byte, int, error) {
	path := strings.TrimPrefix(url, c.BaseURL)
	return c.postMelia(path, payload)
}

type meliaHotelRow struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	City     string  `json:"city"`
	Country  string  `json:"country"`
	Category float64 `json:"category"`
	MinPrice float64 `json:"minPrice"`
	Currency string  `json:"currency"`
	URL      string  `json:"url"`
	Image    string  `json:"image"`
}

type meliaSearchResponse struct {
	Total   int             `json:"total"`
	Hotels  []meliaHotelRow `json:"hotels"`
	HasNext bool            `json:"hasNext"`
}

func decodeMeliaSearch(body []byte) (*meliaSearchResponse, error) {
	var resp meliaSearchResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	if len(resp.Hotels) > 0 || resp.Total > 0 {
		return &resp, nil
	}
	var alt struct {
		Total   int             `json:"totalCount"`
		Hotels  []meliaHotelRow `json:"results"`
		HasNext bool            `json:"hasNext"`
	}
	if err := json.Unmarshal(body, &alt); err == nil && len(alt.Hotels) > 0 {
		resp.Total = alt.Total
		resp.Hotels = alt.Hotels
		resp.HasNext = alt.HasNext
		return &resp, nil
	}
	var wrapped struct {
		Data meliaSearchResponse `json:"data"`
	}
	if err := json.Unmarshal(body, &wrapped); err == nil && len(wrapped.Data.Hotels) > 0 {
		return &wrapped.Data, nil
	}
	return &resp, nil
}

func (r *meliaSearchResponse) toResult(query string, page, pageSize int, brand, base string) *HotelSearchResult {
	hits := make([]HotelHit, 0, len(r.Hotels))
	for _, h := range r.Hotels {
		price := ""
		if h.MinPrice > 0 {
			price = fmt.Sprintf("%.2f", h.MinPrice)
		}
		b := h.Brand
		if brand != "" {
			b = brand
		}
		hits = append(hits, HotelHit{
			ID: h.Code, Name: h.Name, Brand: b, City: h.City, Country: h.Country,
			Stars: h.Category, Price: price, Currency: h.Currency,
			HotelURL: tkbase.Absolutize(base, h.URL), ImageURL: tkbase.Absolutize(base, h.Image),
		})
	}
	return &HotelSearchResult{
		Query: query, Total: r.Total, Page: page, PageSize: pageSize,
		HasNext: r.HasNext, Hotels: hits, Brand: brand, Source: "bff",
	}
}
