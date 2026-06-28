package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.senator.es"

// Client talks to Senator public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Senator Hotels & Resorts",
		"Playa Senator",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "senator"), Brand: brand}
}
