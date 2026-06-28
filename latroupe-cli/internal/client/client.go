package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.latroupe.com"

// Client talks to Latroupe public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Latroupe",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "latroupe"), Brand: brand}
}
