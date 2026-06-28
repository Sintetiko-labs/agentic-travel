package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.wyndhamhotels.com"

// Client talks to Wyndham public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Wyndham Hotels & Resorts",
		"Ramada",
		"Wyndham",
		"Tryp",
		"Dolce by Wyndham",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "wyndham"), Brand: brand}
}
