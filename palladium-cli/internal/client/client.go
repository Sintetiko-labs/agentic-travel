package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.palladiumhotelgroup.com"

// Client talks to Palladium public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Palladium Hotel Group",
		"Ushuaïa Ibiza Beach Hotel",
		"Hard Rock Hotel Ibiza",
		"TRS Hotels",
		"Grand Palladium",
		"Palladium Hotels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "palladium"), Brand: brand}
}
