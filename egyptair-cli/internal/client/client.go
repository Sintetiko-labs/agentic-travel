package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.egyptair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Egyptair",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "egyptair"), Brand: brand}
}
