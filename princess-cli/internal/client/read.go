package client

import (
	"github.com/fbelchi/travelkit/hotel"
)

func (c *Client) Read(idOrURL string) (*HotelView, error) {
	brand := c.Brand
	if brand == "" {
		brand = "Princess Hotels"
	}
	r := &hotel.LDReader{
		BaseURL: c.BaseURL, Brand: brand, FetchHTML: c.FetchHTML,
		Lookup: func(id string) (*HotelHit, error) { return hotel.LookupFromSearch(c.Search, id) },
		URLForID: func(id string) string {
			if hit, err := hotel.LookupFromSearch(c.Search, id); err == nil && hit.HotelURL != "" {
				return hit.HotelURL
			}
			return c.BaseURL + "/es/hoteles"
		},
	}
	return r.Read(idOrURL)
}
