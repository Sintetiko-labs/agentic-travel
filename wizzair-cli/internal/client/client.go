package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://wizzair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Wizz Air",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "wizzair"), Brand: brand}
}
