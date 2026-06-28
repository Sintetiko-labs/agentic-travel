package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.limehome.com"

// Client talks to Limehome public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Limehome",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "limehome"), Brand: brand}
}
