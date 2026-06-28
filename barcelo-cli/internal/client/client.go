package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.barcelo.com"

// Client talks to Barceló public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Barceló Hotel Group",
		"Barceló Hotels & Resorts",
		"Royal Hideaway",
		"Occidental Hotels & Resorts",
		"Allegro Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "barcelo"), Brand: brand}
}
