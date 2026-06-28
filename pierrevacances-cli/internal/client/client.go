package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.pierreetvacances.com"

// Client talks to Pierre & Vacances public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Pierre & Vacances",
		"Pierre & Vacances España",
		"Apartamentos Pierre & Vacances",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "pierrevacances"), Brand: brand}
}
