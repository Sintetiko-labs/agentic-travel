package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.medplaya.com"

// Client talks to MedPlaya public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"MedPlaya",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "medplaya"), Brand: brand}
}
