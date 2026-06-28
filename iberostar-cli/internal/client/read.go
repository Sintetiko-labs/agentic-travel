package client

import (
	"fmt"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
)

// Read returns hotel detail.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	id := strings.TrimSpace(idOrURL)
	payload := map[string]any{
		"query": `query Hotel($id:ID!){hotel(id:$id){id name description address city country stars amenities url}}`,
		"variables": map[string]any{"id": urlPathID(id)},
	}
	var resp struct {
		Data struct {
			Hotel struct {
				ID, Name, Description, Address, City, Country, URL string
				Stars                                              float64
				Amenities                                          []string
			} `json:"hotel"`
		} `json:"data"`
	}
	if err := c.PostJSON(c.BaseURL+"/api/graphql", payload, &resp); err != nil {
		return nil, fmt.Errorf("read %q: %w — try IBEROSTAR_COOKIE", idOrURL, err)
	}
	h := resp.Data.Hotel
	return &HotelView{
		ID: h.ID, Name: h.Name, Brand: c.Brand, Description: h.Description,
		Address: h.Address, City: h.City, Country: h.Country, Stars: h.Stars,
		HotelURL: tkbase.Absolutize(c.BaseURL, h.URL), Amenities: h.Amenities,
	}, nil
}

func urlPathID(id string) string {
	if strings.HasPrefix(id, "http") {
		parts := strings.Split(strings.Trim(id, "/"), "/")
		return parts[len(parts)-1]
	}
	return strings.TrimPrefix(id, "/")
}
