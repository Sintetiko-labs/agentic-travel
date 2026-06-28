package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.aerlingus.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Aer Lingus",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "aerlingus"), Brand: brand}
}
