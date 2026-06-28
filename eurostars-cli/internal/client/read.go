package client

import (
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/hotel"
)

// Read returns hotel detail by id or URL.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	brand := brandOrDefault(c.Brand, "Eurostars Hotel Company")
	r := &hotel.LDReader{
		BaseURL:   c.BaseURL,
		Brand:     brand,
		FetchHTML: c.FetchHTML,
		Lookup: func(id string) (*HotelHit, error) {
			return hotel.LookupFromSearch(c.Search, id)
		},
		URLForID: func(id string) string {
			if res, err := c.Search("", 1, 500); err == nil {
				for _, h := range res.Hotels {
					if strings.EqualFold(h.ID, id) {
						return h.HotelURL
					}
				}
			}
			return tkbase.Absolutize(c.BaseURL, "/"+strings.TrimPrefix(id, "/")+".html")
		},
	}
	return r.Read(idOrURL)
}
