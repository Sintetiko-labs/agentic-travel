package client

import (
	"strings"

	"github.com/fbelchi/travelkit/destination"
	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
)

const hotelsHomePath = "/es"

// Search finds hotels via homepage hotel links (apiweb.cataloniahotels.com for prices).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	html, err := c.FetchHTML(c.BaseURL + hotelsHomePath)
	if err != nil {
		return nil, err
	}
	rows := parse.HotelsFromCataloniaLinks(html, c.BaseURL)
	filtered := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if destination.MatchQuery(query, h.Name, h.URL, h.Address, h.ID) {
			filtered = append(filtered, h)
		}
	}
	if len(filtered) == 0 {
		for _, alias := range destination.Expand(query) {
			path := "/es/hotel/catalonia-" + strings.ToLower(strings.ReplaceAll(alias, " ", "-"))
			html, err := c.FetchHTML(c.BaseURL + path)
			if err != nil {
				continue
			}
			ld := parse.HotelsFromJSONLD(html, c.BaseURL)
			if len(ld) == 0 {
				continue
			}
			h := ld[0]
			slug := strings.TrimPrefix(path, "/es/hotel/")
			filtered = append(filtered, parse.HotelLD{
				ID: slug, Name: h.Name, URL: tkbase.Absolutize(c.BaseURL, path),
				Address: h.Address, Stars: h.Stars,
			})
			break
		}
	}
	return hotelsLDToResult(filtered, query, page, pageSize, c.Brand, c.BaseURL, "html-links"), nil
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
			b = "Catalonia Hotels & Resorts"
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
