package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.bahia-principe.com"

// Client talks to Grupo Piñero public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Fiesta Hotels & Resorts",
		"Grupo Piñero",
		"Bahia Principe",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "pinero"), Brand: brand}
}
