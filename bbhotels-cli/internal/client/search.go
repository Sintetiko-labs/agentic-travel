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
	path := tkhotel.BBHotelsCityPath(query)
	if path == "" {
		return nil, fmt.Errorf("search %q: city listing not mapped (try London, Paris, or Berlin)", query)
	}
	html, err := c.FetchHTML(c.BaseURL + path)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("bbhotels"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	if akamai.IsDenied(403, html) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("bbhotels"))
	}
	rows := parse.HotelsFromBBHotelsCity(html, c.BaseURL)
	rows = tkhotel.FilterByBrand(rows, c.Brand)
	if len(rows) == 0 {
		return nil, fmt.Errorf("search %q: no hotels parsed", query)
	}
	b := c.Brand
	if b == "" {
		b = "B&B Hotels"
	}
	return tkhotel.LDToResult(rows, query, page, pageSize, b, c.BaseURL, "city-listing"), nil
}
