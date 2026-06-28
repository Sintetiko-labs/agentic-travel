package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.petitpalace.com"

// Client talks to Petit Palace public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Petit Palace",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "petitpalace"), Brand: brand}
}
