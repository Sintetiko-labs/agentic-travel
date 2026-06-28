package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.royalairmaroc.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Royal Air Maroc",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "royalairmaroc"), Brand: brand}
}
