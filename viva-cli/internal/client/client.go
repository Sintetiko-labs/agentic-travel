package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.vivahotels.com"

// Client talks to Viva Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Viva Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "viva"), Brand: brand}
}
