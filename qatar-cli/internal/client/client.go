package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.qatarairways.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Qatar Airways",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "qatar"), Brand: brand}
}
