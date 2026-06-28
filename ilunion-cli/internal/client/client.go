package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.ilunionhotels.com"

// Client talks to Ilunion public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Ilunion Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "ilunion"), Brand: brand}
}
