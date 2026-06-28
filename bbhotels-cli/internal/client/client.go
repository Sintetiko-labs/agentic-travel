package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hotel-bb.com"

// Client talks to B&B Hotels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"B&B Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "bbhotels"), Brand: brand}
}
