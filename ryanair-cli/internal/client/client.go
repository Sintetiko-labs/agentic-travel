package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.ryanair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Ryanair",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "ryanair"), Brand: brand}
}
