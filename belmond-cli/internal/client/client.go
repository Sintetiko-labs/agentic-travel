package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.belmond.com"

// Client talks to Belmond public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Belmond",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "belmond"), Brand: brand}
}
