package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.besthotels.es"

// Client talks to Best Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Best Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "besthotels"), Brand: brand}
}
