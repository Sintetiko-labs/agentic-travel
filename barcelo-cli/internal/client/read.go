package client

import (
	"fmt"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/parse"
)

// Read returns hotel detail by slug or URL.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	id := strings.TrimSpace(idOrURL)
	if id == "" {
		return nil, fmt.Errorf("id or url required")
	}
	if !strings.HasPrefix(id, "http") {
		id = tkbase.Absolutize(c.BaseURL, "/content/barcelo/es/es-es/"+strings.TrimPrefix(id, "/"))
	}
	res, err := c.Search("", 1, 500)
	if err != nil {
		return nil, err
	}
	slug := strings.ToLower(strings.TrimPrefix(id, c.BaseURL))
	for _, h := range res.Hotels {
		if strings.EqualFold(h.ID, id) || strings.Contains(strings.ToLower(h.HotelURL), strings.ToLower(slug)) {
			return &HotelView{
				ID:       h.ID,
				Name:     h.Name,
				Brand:    h.Brand,
				City:     h.City,
				Stars:    h.Stars,
				HotelURL: h.HotelURL,
				Price:    PriceInfo{Currency: h.Currency, Price: h.Price},
			}, nil
		}
	}
	html, err := c.FetchHTML(id)
	if err != nil {
		return nil, err
	}
	rows := parse.HotelsFromJSONLD(html, c.BaseURL)
	if len(rows) == 0 {
		return nil, fmt.Errorf("hotel not found at %q", idOrURL)
	}
	h := rows[0]
	return &HotelView{
		ID:       slugFromURL(h.URL),
		Name:     h.Name,
		Brand:    brandFor(c.Brand, h.Name),
		Address:  h.Address,
		Stars:    h.Stars,
		HotelURL: h.URL,
	}, nil
}
