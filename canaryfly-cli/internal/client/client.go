package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.canaryfly.es"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Canaryfly",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "canaryfly"), Brand: brand}
}
