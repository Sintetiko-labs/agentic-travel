package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.mandarinoriental.com"

// Client talks to Mandarin Oriental public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Mandarin Oriental",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "mandarin"), Brand: brand}
}
