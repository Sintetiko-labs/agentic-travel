package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.servigroup.com"

// Client talks to Servigroup public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hoteles Servigroup",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "servigroup"), Brand: brand}
}
