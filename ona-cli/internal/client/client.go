package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.onahotels.com"

// Client talks to Ona Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Ona Hotels",
		"Ona Hotels & Apartments",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "ona"), Brand: brand}
}
