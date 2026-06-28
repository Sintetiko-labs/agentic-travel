package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/destination"
	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search finds hotels via H10 Angular ng-state on destination pages.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)

	paths := destinationPaths(query)
	var rows []parse.H10Hotel
	for _, p := range paths {
		html, err := c.FetchHTML(c.BaseURL + p)
		if err != nil {
			continue
		}
		if found := parse.ParseH10NgState(html, c.BaseURL, query); len(found) > 0 {
			rows = append(rows, found...)
		}
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("search %q: no hotels — try a destination like Barcelona or Madrid", query)
	}
	rows = dedupeH10(rows)
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
		price := ""
		if h.Price > 0 {
			price = fmt.Sprintf("%.2f", h.Price)
		}
		slug := h.Slug
		if slug == "" {
			slug = h.ID
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: brandOrDefault(c.Brand),
			City: h.City, Country: h.Country, Stars: h.Stars,
			Price: price, Currency: h.Currency,
			HotelURL: tkbase.Absolutize(c.BaseURL, "/es/"+slug), ImageURL: h.ImageURL,
		})
	}
	return &HotelSearchResult{
		Query: query, Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Hotels: hits, Brand: c.Brand, Source: "ng-state",
	}, nil
}

func destinationPaths(query string) []string {
	var paths []string
	for _, term := range destination.Expand(query) {
		slug := strings.ToLower(strings.ReplaceAll(term, " ", "-"))
		paths = append(paths,
			"/es/hoteles-"+slug,
			"/es/hoteles/"+slug,
		)
	}
	paths = append(paths, "/es/hoteles/espana")
	return paths
}

func dedupeH10(rows []parse.H10Hotel) []parse.H10Hotel {
	seen := map[string]bool{}
	out := make([]parse.H10Hotel, 0, len(rows))
	for _, h := range rows {
		key := h.ID
		if key == "" {
			key = h.Slug
		}
		if key == "" || seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, h)
	}
	return out
}

func brandOrDefault(brand string) string {
	if brand != "" {
		return brand
	}
	return "H10 Hotels"
}
