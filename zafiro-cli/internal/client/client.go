package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.zafirohotels.com"

// Client talks to Zafiro public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Zafiro Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "zafiro"), Brand: brand}
}
