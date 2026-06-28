package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.h10hotels.com"

// Client talks to H10 public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"H10 Hotels",
		"H10",
		"Ocean by H10",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "h10"), Brand: brand}
}
