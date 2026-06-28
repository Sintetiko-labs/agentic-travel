package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.airarabia.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air Arabia",
		"Air Arabia Maroc",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "airarabia"), Brand: brand}
}
