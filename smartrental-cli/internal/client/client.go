package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.smartrental.com"

// Client talks to SmartRental public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"SmartRental",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "smartrental"), Brand: brand}
}
