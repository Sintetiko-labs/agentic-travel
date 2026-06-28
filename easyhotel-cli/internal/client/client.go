package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.easyhotel.com"

// Client talks to easyHotel public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"easyHotel",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "easyhotel"), Brand: brand}
}
