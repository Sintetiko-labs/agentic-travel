package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.magiccostablanca.com"

// Client talks to Magic Costa Blanca public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Magic Costa Blanca",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "magic"), Brand: brand}
}
