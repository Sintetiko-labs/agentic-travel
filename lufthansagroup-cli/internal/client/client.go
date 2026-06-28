package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.lufthansa.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Lufthansa",
		"Lufthansa City Airlines",
		"Discover Airlines",
		"Swiss",
		"Austrian Airlines",
		"Brussels Airlines",
		"Eurowings",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "lufthansagroup"), Brand: brand}
}
