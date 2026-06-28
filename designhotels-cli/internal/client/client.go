package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.designhotels.com"

// Client talks to Design Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Design Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "designhotels"), Brand: brand}
}
