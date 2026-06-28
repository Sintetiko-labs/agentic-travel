package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hyatt.com"

// Client talks to Hyatt public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hyatt",
		"Grand Hyatt",
		"Hyatt Regency",
		"Hyatt Centric",
		"Thompson Hotels",
		"Andaz",
		"Alua Hotels",
		"Dreams Resorts",
		"Secrets Resorts",
		"Zoëtry",
		"Inclusive Collection",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "hyatt"), Brand: brand}
}
