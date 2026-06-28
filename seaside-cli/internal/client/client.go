package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.seaside-collection.com"

// Client talks to Seaside Collection public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Seaside Collection",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "seaside"), Brand: brand}
}
