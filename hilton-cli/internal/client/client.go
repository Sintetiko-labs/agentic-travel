package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hilton.com"

// Client talks to Hilton public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Hilton",
		"Hilton Hotels & Resorts",
		"Conrad",
		"Waldorf Astoria",
		"DoubleTree by Hilton",
		"Canopy by Hilton",
		"Curio Collection",
		"Hampton by Hilton",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "hilton"), Brand: brand}
}
