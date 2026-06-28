package client

import (
	"fmt"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/parse"
)

// Read returns hotel detail by id or URL.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	id := strings.TrimSpace(idOrURL)
	if id == "" {
		return nil, fmt.Errorf("id or url required")
	}
	if !strings.HasPrefix(id, "http") {
		id = tkbase.Absolutize(c.BaseURL, "/es/hotel/"+strings.TrimPrefix(id, "/"))
	}
	html, err := c.FetchHTML(id)
	if err != nil {
		return nil, err
	}
	rows := parse.ParseRIUNgState(html, c.BaseURL, "")
	if len(rows) == 0 {
		return nil, fmt.Errorf("hotel not found at %q", idOrURL)
	}
	h := rows[0]
	price := ""
	if h.Price > 0 {
		price = fmt.Sprintf("%.2f", h.Price)
	}
	return &HotelView{
		ID:       h.ID,
		Name:     h.Name,
		Brand:    brandOrDefault(c.Brand),
		City:     h.City,
		Country:  h.Country,
		Stars:    h.Stars,
		HotelURL: tkbase.Absolutize(c.BaseURL, "/es/hotel/"+h.Slug),
		Price:    PriceInfo{Price: price, Currency: h.Currency, PerNight: true},
	}, nil
}
