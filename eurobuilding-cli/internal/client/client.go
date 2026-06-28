package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.eurobuilding.es"

// Client talks to Eurobuilding public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Eurobuilding",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "eurobuilding"), Brand: brand}
}
