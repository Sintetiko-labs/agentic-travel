package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.singaporeair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Singapore Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "singaporeairlines"), Brand: brand}
}
