package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.caboverdeairlines.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Cabo Verde Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "caboverde"), Brand: brand}
}
