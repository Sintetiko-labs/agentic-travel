package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.melia.com"

// Client talks to Meliá public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Meliá Hotels International",
		"Meliá",
		"Gran Meliá",
		"ME by Meliá",
		"The Meliá Collection",
		"Paradisus",
		"INNSiDE by Meliá",
		"Sol by Meliá",
		"ZEL",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "melia"), Brand: brand}
}
