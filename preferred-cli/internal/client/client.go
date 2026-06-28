package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.preferredhotels.com"

// Client talks to Preferred Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Preferred Hotels & Resorts",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "preferred"), Brand: brand}
}
