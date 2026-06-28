package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search queries Meliá hotel search BFF (requires Akamai session cookie for live calls).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	payload := map[string]any{
		"text":     query,
		"language": "es",
		"market":   "ES",
		"page":     page,
		"size":     pageSize,
	}
	var resp meliaSearchResponse
	err := c.PostJSON(c.BaseURL+"/services/search/hotels/v2/search", payload, &resp)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("melia"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	return resp.toResult(query, page, pageSize, c.Brand, c.BaseURL), nil
}

type meliaSearchResponse struct {
	Total   int `json:"total"`
	Hotels  []struct {
		Code        string  `json:"code"`
		Name        string  `json:"name"`
		Brand       string  `json:"brand"`
		City        string  `json:"city"`
		Country     string  `json:"country"`
		Category    float64 `json:"category"`
		MinPrice    float64 `json:"minPrice"`
		Currency    string  `json:"currency"`
		URL         string  `json:"url"`
		Image       string  `json:"image"`
	} `json:"hotels"`
	HasNext bool `json:"hasNext"`
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
