package client

import (
	"fmt"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/hotel"
	"github.com/fbelchi/travelkit/parse"
)

func (c *Client) Read(idOrURL string) (*HotelView, error) {
	id := strings.TrimSpace(idOrURL)
	if id == "" {
		return nil, fmt.Errorf("id or url required")
	}
	if hit, err := hotel.LookupFromSearch(c.Search, id); err == nil {
		return &HotelView{
			ID: hit.ID, Name: hit.Name, Brand: brandOrDefault(c.Brand),
			City: hit.City, Country: hit.Country, Stars: hit.Stars, HotelURL: hit.HotelURL,
			Price: PriceInfo{Price: hit.Price, Currency: hit.Currency, PerNight: true},
		}, nil
	}
	url := id
	if !strings.HasPrefix(url, "http") {
		url = tkbase.Absolutize(c.BaseURL, "/es/"+strings.TrimPrefix(id, "/"))
	}
	html, err := c.FetchHTML(url)
	if err != nil {
		return nil, err
	}
	rows := parse.ParseH10NgState(html, c.BaseURL, "")
	if len(rows) == 0 {
		return nil, fmt.Errorf("hotel not found at %q", idOrURL)
	}
	h := rows[0]
	price := ""
	if h.Price > 0 {
		price = fmt.Sprintf("%.2f", h.Price)
	}
	slug := h.Slug
	if slug == "" {
		slug = h.ID
	}
	return &HotelView{
		ID: h.ID, Name: h.Name, Brand: brandOrDefault(c.Brand),
		City: h.City, Country: h.Country, Stars: h.Stars,
		HotelURL: tkbase.Absolutize(c.BaseURL, "/es/"+slug),
		Price:    PriceInfo{Price: price, Currency: h.Currency, PerNight: true},
	}, nil
}
