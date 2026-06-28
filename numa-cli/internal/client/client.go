package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.numastays.com"

// Client talks to Numa public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Numa",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "numa"), Brand: brand}
}
