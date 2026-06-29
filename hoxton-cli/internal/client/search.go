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
	html, err := c.FetchHTML(c.BaseURL + "/")
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("hoxton"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	if akamai.IsDenied(403, html) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("hoxton"))
	}
	rows := parse.HotelsFromHoxtonHome(html, c.BaseURL)
	filtered := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if tkhotel.MatchHoxtonSlug(query, h) {
			filtered = append(filtered, h)
		}
	}
	if len(filtered) == 0 {
		return nil, fmt.Errorf("search %q: no hotels parsed", query)
	}
	b := c.Brand
	if b == "" {
		b = "The Hoxton"
	}
	return tkhotel.LDToResult(filtered, query, page, pageSize, b, c.BaseURL, "homepage"), nil
}
