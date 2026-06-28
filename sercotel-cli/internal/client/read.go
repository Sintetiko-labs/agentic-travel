package client

import (
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/hotel"
)

func (c *Client) Read(idOrURL string) (*HotelView, error) {
	brand := brandOrDefault(c.Brand, "Sercotel")
	r := &hotel.LDReader{
		BaseURL: c.BaseURL, Brand: brand, FetchHTML: c.FetchHTML,
		Lookup: func(id string) (*HotelHit, error) { return hotel.LookupFromSearch(c.Search, id) },
		URLForID: func(id string) string {
			return tkbase.Absolutize(c.BaseURL, "/es/"+strings.TrimPrefix(id, "sercotel-"))
		},
	}
	return r.Read(idOrURL)
}
