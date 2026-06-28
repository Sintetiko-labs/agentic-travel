package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hoteles-poseidon.com"

// Client talks to Hoteles Poseidón public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hoteles Poseidón",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "poseidon"), Brand: brand}
}
