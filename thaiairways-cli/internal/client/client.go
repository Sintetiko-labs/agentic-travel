package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.thaiairways.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Thai Airways",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "thaiairways"), Brand: brand}
}
