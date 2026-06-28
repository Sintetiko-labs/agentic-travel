package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.bintercanarias.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Binter",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "binter"), Brand: brand}
}
