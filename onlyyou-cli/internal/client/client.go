package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.onlyyouhotels.com"

// Client talks to Only YOU public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Only YOU Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "onlyyou"), Brand: brand}
}
