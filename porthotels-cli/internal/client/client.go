package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.porthotels.es"

// Client talks to Port Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Port Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "porthotels"), Brand: brand}
}
