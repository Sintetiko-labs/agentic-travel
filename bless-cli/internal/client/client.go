package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.blesscollectionhotels.com"

// Client talks to BLESS Collection public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"BLESS Collection Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "bless"), Brand: brand}
}
