package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.belivehotels.com"

// Client talks to Be Live public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Be Live Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "belive"), Brand: brand}
}
