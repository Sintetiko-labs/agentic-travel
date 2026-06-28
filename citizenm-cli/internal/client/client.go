package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.citizenm.com"

// Client talks to citizenM public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"citizenM",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "citizenm"), Brand: brand}
}
