package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.tenthotels.com"

// Client talks to Tent Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Tent Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "tent"), Brand: brand}
}
