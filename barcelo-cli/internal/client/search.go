package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/parse"
)

const hotelsListingPath = "/es/hoteles"

// Search finds hotels via JSON-LD embedded in the hotel listing page.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	html, err := c.FetchHTML(c.BaseURL + hotelsListingPath)
	if err != nil {
		return nil, err
	}
	if akamai.IsDenied(403, html) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("barcelo"))
	}
	rows := parse.HotelsFromJSONLD(html, c.BaseURL)
	q := strings.ToLower(query)
	filtered := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if q == "" || strings.Contains(strings.ToLower(h.Name), q) ||
			strings.Contains(strings.ToLower(h.Address), q) {
			filtered = append(filtered, h)
		}
	}
	total := len(filtered)
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}
	hits := make([]HotelHit, 0, end-start)
	for _, h := range filtered[start:end] {
		slug := slugFromURL(h.URL)
		hits = append(hits, HotelHit{
			ID:       slug,
			Name:     h.Name,
			Brand:    brandFor(c.Brand, h.Name),
			City:     query,
			Stars:    h.Stars,
			HotelURL: h.URL,
			ImageURL: h.ImageURL,
		})
	}
	return &HotelSearchResult{
		Query:    query,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		HasNext:  total > page*pageSize,
		Hotels:   hits,
		Brand:    c.Brand,
		Source:   "json-ld",
	}, nil
}

func slugFromURL(raw string) string {
	if raw == "" {
		return ""
	}
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) == 0 {
		return raw
	}
	return parts[len(parts)-1]
}

func brandFor(selected, hotelName string) string {
	if selected != "" {
		return selected
	}
	low := strings.ToLower(hotelName)
	switch {
	case strings.Contains(low, "royal hideaway"):
		return "Royal Hideaway"
	case strings.Contains(low, "occidental"):
		return "Occidental Hotels & Resorts"
	case strings.Contains(low, "allegro"):
		return "Allegro Hotels"
	default:
		return "Barceló Hotels & Resorts"
	}
}
