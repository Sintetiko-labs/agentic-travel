package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
	tkhotel "github.com/fbelchi/travelkit/hotel"
)

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
	return hotelsLDToResult(filtered, query, page, pageSize, c.Brand, "marriott", c.BaseURL, "findHotels"), nil
}

func marriottSearchURL(query string) string {
	city := strings.TrimSpace(query)
	country := marriottCountryCode(query)
	from := time.Now().AddDate(0, 0, 14).Format("01/02/2006")
	to := time.Now().AddDate(0, 0, 15).Format("01/02/2006")
	return fmt.Sprintf("/search/findHotels.mi?destinationAddress.city=%s&destinationAddress.country=%s&roomCount=1&numAdultsPerRoom=2&lengthOfStay=1&fromDate=%s&toDate=%s&deviceType=desktop-web&view=list",
		strings.ReplaceAll(city, " ", "+"), country, from, to)
}

func marriottCountryCode(query string) string {
	switch strings.ToLower(strings.TrimSpace(query)) {
	case "madrid", "barcelona", "valencia", "seville", "sevilla", "malaga", "bilbao":
		return "ES"
	case "paris", "lyon", "marseille":
		return "FR"
	case "berlin", "munich", "münchen", "frankfurt":
		return "DE"
	default:
		return "GB"
	}
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

func hotelsLDToResult(rows []parse.HotelLD, query string, page, pageSize int, brand, parent, base, source string) *HotelSearchResult {
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
		if b == "" && parent != "" {
			b = tkhotel.InferBrand(parent, h.Name)
		}
		if b == "" {
			b = "Marriott"
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: b, City: query,
			Stars: h.Stars, HotelURL: tkbase.Absolutize(base, h.URL), ImageURL: h.ImageURL,
		})
	}
	return &HotelSearchResult{
		Query: query, Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Hotels: hits, Brand: brand, Source: source,
	}
}
