package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hotusa.com"

const SessionStartURL = "https://www.grupohotusa.com/es"

// Client talks to Hotusa public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hotusa",
		"Crisol Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "hotusa"), Brand: brand}
}
