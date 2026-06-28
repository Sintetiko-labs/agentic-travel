package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.st-christophers.co.uk"

// Client talks to St Christopher's public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"St Christopher's Inns Iberia",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "stchristophers"), Brand: brand}
}
