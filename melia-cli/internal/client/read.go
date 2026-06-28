package client

import (
	"fmt"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
)

// Read returns hotel detail by code or URL.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	id := strings.TrimSpace(idOrURL)
	if id == "" {
		return nil, fmt.Errorf("id or url required")
	}
	path := "/services/content/hotels/v1/hotel/" + strings.TrimPrefix(id, "/")
	if strings.HasPrefix(id, "http") {
		path = strings.TrimPrefix(id, c.BaseURL)
	}
	var resp struct {
		Code        string   `json:"code"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Address     string   `json:"address"`
		City        string   `json:"city"`
		Country     string   `json:"country"`
		Category    float64  `json:"category"`
		Amenities   []string `json:"amenities"`
		URL         string   `json:"url"`
	}
	if err := c.GetJSON(path, &resp); err != nil {
		return nil, fmt.Errorf("read %q: %w — try MELIA_COOKIE", idOrURL, err)
	}
	return &HotelView{
		ID: resp.Code, Name: resp.Name, Brand: c.Brand, Description: resp.Description,
		Address: resp.Address, City: resp.City, Country: resp.Country, Stars: resp.Category,
		HotelURL: tkbase.Absolutize(c.BaseURL, resp.URL), Amenities: resp.Amenities,
	}, nil
}
