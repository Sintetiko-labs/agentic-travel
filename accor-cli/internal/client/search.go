package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
	tkhotel "github.com/fbelchi/travelkit/hotel"
	"github.com/fbelchi/travelkit/parse"
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
	path := tkhotel.AccorDestinationPath(query)
	source := "destination"
	html, err := c.FetchHTML(c.BaseURL + path)
	if err != nil || akamai.IsDenied(403, html) || len(parse.HotelsFromAccorDestination(html, c.BaseURL)) == 0 {
		fb := tkhotel.AccorSearchFallbackPath(query)
		if fb == "" {
			if err != nil {
				if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
					return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("accor"))
				}
				return nil, fmt.Errorf("search %q: %w", query, err)
			}
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("accor"))
		}
		source = "search-fallback"
		html, err = c.FetchHTML(c.BaseURL + fb)
		if err != nil {
			if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
				return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("accor"))
			}
			return nil, fmt.Errorf("search %q: %w", query, err)
		}
	}
	if akamai.IsDenied(403, html) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("accor"))
	}
	rows := parse.HotelsFromAccorDestination(html, c.BaseURL)
	rows = tkhotel.FilterHotelLD(rows, query)
	rows = tkhotel.FilterByBrand(rows, c.Brand)
	if len(rows) == 0 {
		return nil, fmt.Errorf("search %q: no hotels parsed", query)
	}
	return tkhotel.LDToResultParent(rows, query, page, pageSize, c.Brand, "accor", c.BaseURL, source), nil
}
