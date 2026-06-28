package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.easyjet.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"easyJet",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "easyjet"), Brand: brand}
}
