package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.globales.com"

// Client talks to Globales public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Globales Hotels",
		"Hoteles Globales",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "globales"), Brand: brand}
}
