package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.umusichotels.com"

// Client talks to UMusic Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"UMusic Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "umusic"), Brand: brand}
}
