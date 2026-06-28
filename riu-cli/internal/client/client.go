package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.riu.com"

// Client talks to RIU public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"RIU Hotels & Resorts",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "riu"), Brand: brand}
}
