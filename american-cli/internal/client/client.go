package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.aa.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"American Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "american"), Brand: brand}
}
