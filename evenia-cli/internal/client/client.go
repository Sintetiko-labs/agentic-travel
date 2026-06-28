package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.eveniahotels.com"

// Client talks to Evenia public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Evenia Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "evenia"), Brand: brand}
}
