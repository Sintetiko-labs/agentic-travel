package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.parador.es"

// Client talks to Paradores public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Paradores",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "paradores"), Brand: brand}
}
