package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.emirates.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Emirates",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "emirates"), Brand: brand}
}
