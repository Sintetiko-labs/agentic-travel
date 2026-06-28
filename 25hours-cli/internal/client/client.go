package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.25hours-hotels.com"

// Client talks to 25hours Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"25hours Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "25hours"), Brand: brand}
}
