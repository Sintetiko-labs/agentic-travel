package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.guitarthotels.com"

// Client talks to Guitart public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Guitart Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "guitart"), Brand: brand}
}
