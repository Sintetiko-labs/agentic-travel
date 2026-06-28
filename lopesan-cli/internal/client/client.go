package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.lopesan.com"

// Client talks to Lopesan public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Lopesan Hotel Group",
		"Abora by Lopesan",
		"Lopesan Hotels",
		"Lopesan Collection",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "lopesan"), Brand: brand}
}
