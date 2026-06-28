package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.smartwings.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Smartwings",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "smartwings"), Brand: brand}
}
