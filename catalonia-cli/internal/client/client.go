package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.cataloniahotels.com"

// Client talks to Catalonia Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Catalonia Hotels & Resorts",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "catalonia"), Brand: brand}
}
