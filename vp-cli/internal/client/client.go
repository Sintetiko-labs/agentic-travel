package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.vp-hoteles.com"

// Client talks to VP Hoteles public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"VP Hoteles",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "vp"), Brand: brand}
}
