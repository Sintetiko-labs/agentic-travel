package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.privilegestyle.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Privilege Style",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "privilegestyle"), Brand: brand}
}
