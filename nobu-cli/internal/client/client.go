package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.nobuhotels.com"

// Client talks to Nobu Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Nobu Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "nobu"), Brand: brand}
}
