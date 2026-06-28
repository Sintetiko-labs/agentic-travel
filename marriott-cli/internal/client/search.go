package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search queries Marriott findHotels (Akamai session required from residential IP).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, fmt.Errorf("destination required")
	}
	path := marriottSearchURL(query)
	html, err := c.fetchSearchHTML(path)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, marriottBlockedErr(c.Cookie)
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	if akamai.IsDenied(403, html) {
		return nil, marriottBlockedErr(c.Cookie)
	}
	rows := parse.HotelsFromMarriottSearch(html, c.BaseURL)
	if len(rows) == 0 {
		return nil, fmt.Errorf("search %q: no hotels parsed — run: marriott session chrome --wait", query)
	}
	filtered := filterByBrand(rows, c.Brand)
	return hotelsLDToResult(filtered, query, page, pageSize, brandFor(c.Brand), c.BaseURL, "findHotels"), nil
}

func marriottSearchURL(query string) string {
	from := time.Now().AddDate(0, 0, 14)
	to := from.AddDate(0, 0, 1)
	return marriottSearchURLWithDates(query, from, to)
}

func marriottBlockedErr(cookie string) error {
	if akamai.SessionReady(cookie) {
		return fmt.Errorf("akamai blocked — saved cookies may be stale or not valid for CLI requests; re-run: marriott session chrome --wait --timeout 3m")
	}
	return fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("marriott"))
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
			b = "Marriott"
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
	return "Marriott"
}
