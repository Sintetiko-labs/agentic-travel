package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.sercotelhoteles.com"

// Client talks to Sercotel public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Sercotel",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "sercotel"), Brand: brand}
}
