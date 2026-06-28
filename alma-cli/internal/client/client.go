package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.almahotels.com"

// Client talks to Alma Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Alma Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "alma"), Brand: brand}
}
