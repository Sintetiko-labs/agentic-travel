package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.bestwestern.com"

// Client talks to Best Western public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Best Western",
		"Best Western Plus",
		"Best Western Premier",
		"BWH Hotel Group",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "bestwestern"), Brand: brand}
}
