package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.radissonhotels.com"

// Client talks to Radisson public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Radisson Hotel Group",
		"Radisson Blu",
		"Radisson RED",
		"Radisson Collection",
		"Park Inn by Radisson",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "radisson"), Brand: brand}
}
