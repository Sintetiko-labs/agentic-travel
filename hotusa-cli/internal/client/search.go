package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
)

func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)

	var rows []parse.HotelLD
	seen := map[string]bool{}
	var lastErr error
	for _, p := range hotusaPaths(query) {
		html, err := c.fetchHotusaHTML(p)
		if err != nil {
			lastErr = err
			continue
		}
		for _, h := range parse.HotelsFromHotusaLinks(html, c.BaseURL) {
			if seen[h.URL] {
				continue
			}
			seen[h.URL] = true
			rows = append(rows, h)
		}
		if len(rows) > 0 {
			break
		}
	}
	if len(rows) == 0 {
		if lastErr != nil {
			return nil, fmt.Errorf("search %q: %w", query, hotusaSearchErr(lastErr))
		}
		return nil, fmt.Errorf("search %q: no hotels found — run: hotusa session chrome --wait", query)
	}
	filtered := filterHotelLD(rows, query)
	if len(filtered) == 0 {
		return nil, fmt.Errorf("search %q: no hotels matched", query)
	}
	return ldToResult(filtered, query, page, pageSize, brandOrDefault(c.Brand, "Hotusa"), c.BaseURL, "html-links"), nil
}

func filterHotelLD(rows []parse.HotelLD, query string) []parse.HotelLD {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return rows
	}
	out := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if strings.Contains(strings.ToLower(h.Name), q) ||
			strings.Contains(strings.ToLower(h.URL), q) ||
			strings.Contains(strings.ToLower(h.Address), q) ||
			strings.Contains(strings.ToLower(h.ID), q) {
			out = append(out, h)
		}
	}
	return out
}

func ldToResult(rows []parse.HotelLD, query string, page, pageSize int, brand, base, source string) *HotelSearchResult {
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
		city := query
		if h.Address != "" {
			city = h.Address
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: brand, City: city,
			Stars: h.Stars, HotelURL: tkbase.Absolutize(base, h.URL), ImageURL: h.ImageURL,
		})
	}
	return &HotelSearchResult{
		Query: query, Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Hotels: hits, Brand: brand, Source: source,
	}
}

func brandOrDefault(brand, fallback string) string {
	if brand != "" {
		return brand
	}
	return fallback
}
