package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.mamashelter.com"

// Client talks to Mama Shelter public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Mama Shelter",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "mamashelter"), Brand: brand}
}
