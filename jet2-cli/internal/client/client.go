package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.jet2.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Jet2",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "jet2"), Brand: brand}
}
