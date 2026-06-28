package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.rosewoodhotels.com"

// Client talks to Rosewood public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Rosewood",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "rosewood"), Brand: brand}
}
