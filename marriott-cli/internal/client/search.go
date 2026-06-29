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
	html, err := c.FetchHTML(c.BaseURL + path)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("marriott"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	if akamai.IsDenied(403, html) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("marriott"))
	}
	rows := parse.HotelsFromMarriottSearch(html, c.BaseURL)
	if len(rows) == 0 {
		return nil, fmt.Errorf("search %q: no hotels parsed — run: marriott session chrome --wait", query)
	}
	filtered := filterByBrand(rows, c.Brand)
	return hotelsLDToResult(filtered, query, page, pageSize, brandFor(c.Brand), c.BaseURL, "findHotels"), nil
}

func marriottSearchURL(query string) string {
	city := strings.TrimSpace(query)
	from := time.Now().AddDate(0, 0, 14).Format("01/02/2006")
	to := time.Now().AddDate(0, 0, 15).Format("01/02/2006")
	return fmt.Sprintf("/search/findHotels.mi?destinationAddress.city=%s&destinationAddress.country=GB&roomCount=1&numAdultsPerRoom=2&lengthOfStay=1&fromDate=%s&toDate=%s&deviceType=desktop-web&view=list",
		strings.ReplaceAll(city, " ", "+"), from, to)
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
