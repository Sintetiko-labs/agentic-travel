package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.oneshothotels.com"

// Client talks to One Shot public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"One Shot Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "oneshot"), Brand: brand}
}
