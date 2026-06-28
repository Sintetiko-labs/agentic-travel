package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.safestay.com"

// Client talks to Safestay public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Safestay",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "safestay"), Brand: brand}
}
