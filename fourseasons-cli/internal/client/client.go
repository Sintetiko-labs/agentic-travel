package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.fourseasons.com"

// Client talks to Four Seasons public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Four Seasons",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "fourseasons"), Brand: brand}
}
