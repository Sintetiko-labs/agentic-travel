package client

import (
	"fmt"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/hotel"
)

func (c *Client) Read(idOrURL string) (*HotelView, error) {
	id := strings.TrimSpace(idOrURL)
	if id == "" {
		return nil, fmt.Errorf("id or url required")
	}
	r := &hotel.LDReader{
		BaseURL: c.BaseURL, Brand: brandOrDefault(c.Brand), FetchHTML: c.FetchHTML,
		URLForID: func(s string) string {
			return tkbase.Absolutize(c.BaseURL, "/es/hotel/"+strings.TrimPrefix(s, "/"))
		},
		Lookup: func(s string) (*HotelHit, error) {
			return hotel.LookupFromSearch(func(q string, p, l int) (*HotelSearchResult, error) {
				return c.Search(q, p, l)
			}, s)
		},
	}
	return r.Read(id)
}
