package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/parse"
)

const hotelsListingPath = "/es/hoteles"

// Search finds hotels via Lopesan hotel detail links on listing pages.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)

	paths := []string{hotelsListingPath, "/es/hoteles/espana"}
	slug := strings.ToLower(strings.ReplaceAll(query, " ", "-"))
	if slug != "" {
		paths = append([]string{"/es/hoteles/espana/" + slug}, paths...)
	}
	var rows []parse.HotelLD
	seen := map[string]bool{}
	for _, p := range paths {
		html, err := c.FetchHTML(c.BaseURL + p)
		if err != nil {
			continue
		}
		for _, h := range parse.HotelsFromLopesanLinks(html, c.BaseURL) {
			if seen[h.URL] {
				continue
			}
			seen[h.URL] = true
			rows = append(rows, h)
		}
	}
	q := strings.ToLower(query)
	filtered := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if q == "" || strings.Contains(strings.ToLower(h.Name), q) ||
			strings.Contains(strings.ToLower(h.URL), q) {
			filtered = append(filtered, h)
		}
	}
	if len(filtered) == 0 {
		return nil, fmt.Errorf("search %q: no hotels found", query)
	}
	return lopesanToResult(filtered, query, page, pageSize, c.Brand), nil
}

func lopesanToResult(rows []parse.HotelLD, query string, page, pageSize int, brand string) *HotelSearchResult {
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
		b := brand
		if b == "" {
			b = "Lopesan Hotel Group"
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: b, City: query,
			HotelURL: h.URL,
		})
	}
	return &HotelSearchResult{
		Query: query, Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Hotels: hits, Brand: brand, Source: "html-links",
	}
}
