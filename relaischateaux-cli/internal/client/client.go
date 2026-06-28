package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.relaischateaux.com"

// Client talks to Relais & Châteaux public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Relais & Châteaux",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "relaischateaux"), Brand: brand}
}
