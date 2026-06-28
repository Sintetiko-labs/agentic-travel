package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.casualhoteles.com"

// Client talks to Casual Hoteles public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Casual Hoteles",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "casual"), Brand: brand}
}
