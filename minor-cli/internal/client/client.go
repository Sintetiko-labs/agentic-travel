package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.minorhotels.com"

// Client talks to Minor Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Avani",
		"Tivoli",
		"Minor Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "minor"), Brand: brand}
}
