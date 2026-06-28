package client

import (
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/hotel"
)

func (c *Client) Read(idOrURL string) (*HotelView, error) {
	brand := c.Brand
	if brand == "" {
		brand = "Catalonia Hotels & Resorts"
	}
	r := &hotel.LDReader{
		BaseURL: c.BaseURL, Brand: brand, FetchHTML: c.FetchHTML,
		URLForID: func(id string) string {
			return tkbase.Absolutize(c.BaseURL, "/es/hotel/"+strings.TrimPrefix(id, "/"))
		},
		Lookup: func(id string) (*HotelHit, error) { return hotel.LookupFromSearch(c.Search, id) },
	}
	return r.Read(idOrURL)
}
