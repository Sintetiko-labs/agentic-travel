package client

import (
	"encoding/json"
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
	body, status, err := c.PostRaw(c.BaseURL+"/services/search/hotels/v2/search", payload)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		if akamai.IsDenied(status, string(body)) {
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

func (c *Client) PostRaw(url string, payload any) ([]byte, int, error) {
	c.Throttle()
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}
	req, err := httpNewPost(url, b)
	if err != nil {
		return nil, 0, err
	}
	c.SetAPIHeaders(req)
	req.Header.Set("content-type", "application/json")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := ioReadAll(resp.Body)
	return body, resp.StatusCode, nil
}

type meliaHotelRow struct {
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
}

type meliaSearchResponse struct {
	Total   int `json:"total"`
	Hotels  []meliaHotelRow `json:"hotels"`
	HasNext bool `json:"hasNext"`
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
		Hotels  []meliaHotelRow   `json:"results"`
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
