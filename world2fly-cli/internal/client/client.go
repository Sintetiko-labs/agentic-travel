package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.world2fly.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"World2Fly",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "world2fly"), Brand: brand}
}
