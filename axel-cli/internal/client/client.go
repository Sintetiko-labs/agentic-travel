package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.axelhotels.com"

// Client talks to Axel Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Axel Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "axel"), Brand: brand}
}
