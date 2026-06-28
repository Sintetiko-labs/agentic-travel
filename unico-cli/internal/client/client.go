package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.unicohotels.com"

// Client talks to Único Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Único Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "unico"), Brand: brand}
}
