package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.sonder.com"

// Client talks to Sonder public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Sonder",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "sonder"), Brand: brand}
}
