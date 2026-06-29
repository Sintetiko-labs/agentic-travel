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
	slug := tkhotel.LimehomeDestinationSlug(query)
	path := "/en/destinations/" + slug + "/"
	source := "destinations"
	html, err := c.FetchHTML(c.BaseURL + path)
	rows := []parse.HotelLD{}
	if err == nil {
		rows = parse.HotelsFromLimehomeDestination(html, c.BaseURL)
	}
	if err != nil || len(rows) == 0 {
		source = "locations"
		html, err = c.FetchHTML(c.BaseURL + "/en/locations/")
		if err != nil {
			if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
				return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("limehome"))
			}
			return nil, fmt.Errorf("search %q: %w", query, err)
		}
		rows = parse.HotelsFromLimehomeDestination(html, c.BaseURL)
	}
	if akamai.IsDenied(403, html) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("limehome"))
	}
	rows = tkhotel.FilterHotelLD(rows, query)
	if len(rows) == 0 {
		return nil, fmt.Errorf("search %q: no hotels parsed", query)
	}
	b := c.Brand
	if b == "" {
		b = "Limehome"
	}
	return tkhotel.LDToResult(rows, query, page, pageSize, b, c.BaseURL, source), nil
}
