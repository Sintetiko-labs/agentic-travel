package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.liberehospitality.com"

// Client talks to Líbere public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Líbere Hospitality",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "libere"), Brand: brand}
}
