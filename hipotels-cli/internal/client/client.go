package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hipotels.com"

// Client talks to Hipotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hipotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "hipotels"), Brand: brand}
}
