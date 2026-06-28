package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.grupotel.com"

// Client talks to Grupotel public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Grupotel",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "grupotel"), Brand: brand}
}
