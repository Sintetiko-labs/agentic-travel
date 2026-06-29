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
	path := parse.BestWesternSearchURL(c.BaseURL, tkhotel.BestWesternPlace(query))
	html, err := c.FetchHTML(path)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("bestwestern"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	if akamai.IsDenied(403, html) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("bestwestern"))
	}
	rows := parse.HotelsFromJSONLD(html, c.BaseURL)
	rows = tkhotel.FilterByBrand(rows, c.Brand)
	if len(rows) == 0 {
		return nil, fmt.Errorf("search %q: no hotels parsed", query)
	}
	return tkhotel.LDToResultParent(rows, query, page, pageSize, c.Brand, "bestwestern", c.BaseURL, "hotel-search"), nil
}
