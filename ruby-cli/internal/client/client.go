package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.ruby-hotels.com"

// Client talks to Ruby Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Ruby Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "ruby"), Brand: brand}
}
