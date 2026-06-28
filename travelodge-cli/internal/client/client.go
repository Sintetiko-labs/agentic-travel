package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.travelodge.co.uk"

// Client talks to Travelodge public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Travelodge",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "travelodge"), Brand: brand}
}
