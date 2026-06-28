package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.vinccihoteles.com"

// Client talks to Vincci public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Vincci Hoteles",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "vincci"), Brand: brand}
}
