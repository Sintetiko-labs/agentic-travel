package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.sohohoteles.com"

// Client talks to Soho Boutique public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Soho Boutique Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "soho"), Brand: brand}
}
