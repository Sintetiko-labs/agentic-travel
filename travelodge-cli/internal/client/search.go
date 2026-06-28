package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
)

const travelodgeSitemapURL = "https://www.travelodge.co.uk/sitemap-fusion.xml"

// Search finds UK hotels via the public sitemap (filters by destination query).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	body, status, err := c.GetRaw(travelodgeSitemapURL)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		if akamai.IsDenied(status, string(body)) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("travelodge"))
		}
		return nil, fmt.Errorf("search %q: HTTP %d: %s", query, status, tkbase.Truncate(string(body), 200))
	}
	rows := parse.HotelsFromTravelodgeSitemap(string(body), query)
	return hotelsLDToResult(rows, query, page, pageSize, brandFor(c.Brand), c.BaseURL, "sitemap"), nil
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
			b = "Travelodge"
		}
		city := query
		if h.Address != "" {
			city = h.Address
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: b, City: city, Country: "GB",
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
	return "Travelodge"
}

func brandOrDefault(selected string) string {
	return brandFor(selected)
}
