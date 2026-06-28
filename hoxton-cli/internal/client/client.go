package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.thehoxton.com"

// Client talks to The Hoxton public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"The Hoxton",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "hoxton"), Brand: brand}
}
