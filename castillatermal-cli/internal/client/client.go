package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.castillatermal.com"

// Client talks to Castilla Termal public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Castilla Termal",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "castillatermal"), Brand: brand}
}
