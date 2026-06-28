package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.lot.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"LOT Polish Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "lot"), Brand: brand}
}
