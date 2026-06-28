package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.aman.com"

// Client talks to Aman public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Aman",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "aman"), Brand: brand}
}
