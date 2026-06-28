package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.slh.com"

// Client talks to Small Luxury Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Small Luxury Hotels of the World",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "slh"), Brand: brand}
}
