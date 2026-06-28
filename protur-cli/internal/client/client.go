package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.protur-hotels.com"

// Client talks to Protur public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Protur Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "protur"), Brand: brand}
}
