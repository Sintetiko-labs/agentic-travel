package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.flylevel.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Level",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "level"), Brand: brand}
}
