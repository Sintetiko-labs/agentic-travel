package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.delta.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Delta Air Lines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "delta"), Brand: brand}
}
