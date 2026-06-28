package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.eurostarshotels.com"

// Client talks to Eurostars public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Eurostars Hotel Company",
		"Eurostars Hotels",
		"Exe Hotels",
		"Ikonik Hotels",
		"Áurea Hotels",
		"Tandem Suites",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "eurostars"), Brand: brand}
}
