package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hoteleselba.com"

// Client talks to Hoteles Elba public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hoteles Elba",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "elba"), Brand: brand}
}
