package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.copaair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Copa Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "copa"), Brand: brand}
}
