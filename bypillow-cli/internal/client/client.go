package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.bypillow.com"

// Client talks to ByPillow public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"ByPillow",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "bypillow"), Brand: brand}
}
