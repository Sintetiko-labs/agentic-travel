package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.gardenhotels.com"

// Client talks to Garden Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Garden Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "garden"), Brand: brand}
}
