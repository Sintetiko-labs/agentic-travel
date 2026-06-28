package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.iberostar.com"

// Client talks to Iberostar public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Iberostar",
		"Iberostar Selection",
		"Iberostar Grand",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "iberostar"), Brand: brand}
}
