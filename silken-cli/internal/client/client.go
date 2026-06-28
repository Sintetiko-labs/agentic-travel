package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hoteles-silken.com"

// Client talks to Silken public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Silken Hoteles",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "silken"), Brand: brand}
}
