package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.derbyhotels.com"

// Client talks to Derby Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Derby Hotels Collection",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "derby"), Brand: brand}
}
