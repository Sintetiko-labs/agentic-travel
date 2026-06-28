package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.nh-hotels.com"

// Client talks to NH Hotel Group public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"NH Hotel Group",
		"NH Hotels",
		"NH Collection",
		"nhow",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "nh"), Brand: brand}
}
