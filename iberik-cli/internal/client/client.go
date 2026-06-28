package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.iberikhotels.com"

// Client talks to Iberik public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Iberik Hoteles",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "iberik"), Brand: brand}
}
