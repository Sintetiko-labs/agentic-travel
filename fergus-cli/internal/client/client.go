package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.fergushotels.com"

// Client talks to Fergus public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Fergus Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "fergus"), Brand: brand}
}
