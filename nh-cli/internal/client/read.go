package client

import (
	"fmt"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
)

// Read returns hotel detail.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	id := strings.TrimSpace(idOrURL)
	if id == "" {
		return nil, fmt.Errorf("id or url required")
	}
	path := "/nh/es/api/v1/hotels/" + urlPathID(id)
	var resp struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Address     string   `json:"address"`
		City        string   `json:"city"`
		Country     string   `json:"country"`
		Stars       float64  `json:"stars"`
		Amenities   []string `json:"amenities"`
		Slug        string   `json:"slug"`
	}
	if err := c.GetJSON(path, &resp); err != nil {
		return nil, fmt.Errorf("read %q: %w — try NH_COOKIE", idOrURL, err)
	}
	return &HotelView{
		ID: resp.ID, Name: resp.Name, Brand: c.Brand, Description: resp.Description,
		Address: resp.Address, City: resp.City, Country: resp.Country, Stars: resp.Stars,
		HotelURL: tkbase.Absolutize(c.BaseURL, "/es/hotel/"+resp.Slug), Amenities: resp.Amenities,
	}, nil
}

func urlPathID(id string) string {
	if strings.HasPrefix(id, "http") {
		parts := strings.Split(strings.Trim(id, "/"), "/")
		return parts[len(parts)-1]
	}
	return strings.TrimPrefix(id, "/")
}
