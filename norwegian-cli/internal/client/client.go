package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.norwegian.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Norwegian",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "norwegian"), Brand: brand}
}
