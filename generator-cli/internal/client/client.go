package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.staygenerator.com"

// Client talks to Generator Hostels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Generator Hostels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "generator"), Brand: brand}
}
