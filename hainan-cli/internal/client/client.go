package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.hainanairlines.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Hainan Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "hainan"), Brand: brand}
}
