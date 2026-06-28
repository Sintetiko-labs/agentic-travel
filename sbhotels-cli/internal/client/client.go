package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.sbhotels.com"

// Client talks to SB Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"SB Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "sbhotels"), Brand: brand}
}
