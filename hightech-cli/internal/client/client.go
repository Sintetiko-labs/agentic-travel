package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hthoteles.com"

// Client talks to High Tech Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"High Tech Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "hightech"), Brand: brand}
}
