package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hospes.com"

// Client talks to Hospes public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hospes Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "hospes"), Brand: brand}
}
