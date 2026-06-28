package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.zenithoteles.com"

// Client talks to Zenit public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Zenit Hoteles",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "zenit"), Brand: brand}
}
