package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.princess-hotels.com"

// Client talks to Princess Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Princess Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "princess"), Brand: brand}
}
