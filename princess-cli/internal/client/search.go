package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/parse"
)

// Search finds hotels via destination page headings on princess-hotels.com.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)

	slug := strings.ToLower(strings.ReplaceAll(query, " ", "-"))
	paths := []string{
		"/es/hoteles/" + slug,
		"/es/hoteles",
	}
	var rows []parse.HotelLD
	for _, p := range paths {
		pageURL := c.BaseURL + p
		html, err := c.FetchHTML(pageURL)
		if err != nil {
			continue
		}
		rows = append(rows, parse.HotelsFromPrincessHeadings(html, c.BaseURL, pageURL)...)
	}
	q := strings.ToLower(query)
	filtered := make([]parse.HotelLD, 0, len(rows))
	seen := map[string]bool{}
	for _, h := range rows {
		if seen[h.ID] {
			continue
		}
		seen[h.ID] = true
		if q != "" && !strings.Contains(strings.ToLower(h.Name), q) &&
			!strings.Contains(strings.ToLower(slug), q) && slug != "" {
			continue
		}
		filtered = append(filtered, h)
	}
	if len(filtered) == 0 {
		return nil, fmt.Errorf("search %q: no hotels — try a destination like Tenerife", query)
	}
	return princessToResult(filtered, query, page, pageSize, c.Brand), nil
}

func princessToResult(rows []parse.HotelLD, query string, page, pageSize int, brand string) *HotelSearchResult {
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
			b = "Princess Hotels"
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: b, City: query, HotelURL: h.URL,
		})
	}
	return &HotelSearchResult{
		Query: query, Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Hotels: hits, Brand: brand, Source: "html-headings",
	}
}
