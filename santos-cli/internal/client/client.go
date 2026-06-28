package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hoteles-santos.com"

// Client talks to Hoteles Santos public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hoteles Santos",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "santos"), Brand: brand}
}
