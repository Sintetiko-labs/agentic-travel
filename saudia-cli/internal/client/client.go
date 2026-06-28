package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.saudia.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Saudia",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "saudia"), Brand: brand}
}
