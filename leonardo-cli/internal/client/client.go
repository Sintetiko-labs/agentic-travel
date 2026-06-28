package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.leonardo-hotels.com"

// Client talks to Leonardo Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Leonardo Hotels",
		"NYX Hotels",
		"Leonardo Royal",
		"Leonardo Boutique",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "leonardo"), Brand: brand}
}
