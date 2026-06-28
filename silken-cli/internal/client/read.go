package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/hotel"
)

// Read returns hotel detail by id or URL.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	brand := brandOrDefault(c.Brand, "Silken Hoteles")
	r := &hotel.LDReader{
		BaseURL:   c.BaseURL,
		Brand:     brand,
		FetchHTML: c.FetchHTML,
		Lookup: func(id string) (*HotelHit, error) {
			return hotel.LookupFromSearch(c.Search, id)
		},
		URLForID: func(id string) string {
			return tkbase.Absolutize(c.BaseURL, "/es/hoteles")
		},
	}
	return r.Read(idOrURL)
}
