package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.coolrooms.com"

// Client talks to CoolRooms public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"CoolRooms Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "coolrooms"), Brand: brand}
}
