package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.icelandair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Icelandair",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "icelandair"), Brand: brand}
}
