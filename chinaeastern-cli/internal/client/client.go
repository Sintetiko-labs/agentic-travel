package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.ceair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"China Eastern",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "chinaeastern"), Brand: brand}
}
