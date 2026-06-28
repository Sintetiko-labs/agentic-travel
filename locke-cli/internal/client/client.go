package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.lockeliving.com"

// Client talks to Locke public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Locke",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "locke"), Brand: brand}
}
