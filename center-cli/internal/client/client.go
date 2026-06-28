package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hotelescenter.com"

// Client talks to Hoteles Center public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hoteles Center",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "center"), Brand: brand}
}
