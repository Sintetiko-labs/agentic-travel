package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.room-matehotels.com"

// Client talks to Room Mate public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Room Mate Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "roommate"), Brand: brand}
}
