package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.abbahoteles.com"

// Client talks to Abba Hoteles public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Abba Hoteles",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "abba"), Brand: brand}
}
