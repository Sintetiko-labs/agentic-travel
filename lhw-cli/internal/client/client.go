package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.lhw.com"

// Client talks to Leading Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Leading Hotels of the World",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "lhw"), Brand: brand}
}
