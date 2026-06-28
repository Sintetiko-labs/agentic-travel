package client

import (
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/hotel"
)

// Read returns hotel detail by id or URL.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	brand := c.Brand
	if brand == "" {
		brand = "Palladium Hotel Group"
	}
	r := &hotel.LDReader{
		BaseURL:   c.BaseURL,
		Brand:     brand,
		FetchHTML: c.FetchHTML,
		Lookup: func(id string) (*HotelHit, error) {
			return hotel.LookupFromSearch(c.Search, id)
		},
		URLForID: func(id string) string {
			if hit, err := hotel.LookupFromSearch(c.Search, id); err == nil {
				return hit.HotelURL
			}
			return tkbase.Absolutize(c.BaseURL, "/es/hoteles/espana/"+strings.TrimPrefix(id, "/"))
		},
	}
	return r.Read(idOrURL)
}
