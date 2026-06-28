package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.alegriahotels.com"

// Client talks to Alegria public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Alegria Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "alegria"), Brand: brand}
}
