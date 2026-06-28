package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.virginhotels.com"

// Client talks to Virgin Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Virgin Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "virgin"), Brand: brand}
}
