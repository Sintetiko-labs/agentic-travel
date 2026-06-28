package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search finds hotels via Hilton destination listing pages (UK: /en/locations/united-kingdom/{city}/).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	path := hiltonLocationPath(query)
	html, err := c.FetchHTML(c.BaseURL + path)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("hilton"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	rows := parse.HotelsFromHiltonLocations(html, c.BaseURL)
	if len(rows) == 0 {
		if akamai.IsDenied(403, html) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("hilton"))
		}
	}
	filtered := filterByBrand(rows, c.Brand)
	return hotelsLDToResult(filtered, query, page, pageSize, brandFor(c.Brand), c.BaseURL, "locations"), nil
}

func hiltonLocationPath(query string) string {
	slug := strings.ToLower(strings.TrimSpace(query))
	slug = strings.ReplaceAll(slug, " ", "-")
	switch slug {
	case "uk", "united-kingdom", "england":
		return "/en/locations/united-kingdom/"
	}
	return "/en/locations/united-kingdom/" + slug + "/"
}

func filterByBrand(rows []parse.HotelLD, brand string) []parse.HotelLD {
	b := strings.TrimSpace(brand)
	if b == "" {
		return rows
	}
	low := strings.ToLower(b)
	out := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if strings.Contains(strings.ToLower(h.Name), low) {
			out = append(out, h)
		}
	}
	return out
}

func hotelsLDToResult(rows []parse.HotelLD, query string, page, pageSize int, brand, base, source string) *HotelSearchResult {
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
			b = "Hilton"
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: b, City: query, Country: "GB",
			Stars: h.Stars, HotelURL: tkbase.Absolutize(base, h.URL), ImageURL: h.ImageURL,
		})
	}
	return &HotelSearchResult{
		Query: query, Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Hotels: hits, Brand: brand, Source: source,
	}
}

func brandFor(selected string) string {
	if selected != "" {
		return selected
	}
	return "Hilton"
}
