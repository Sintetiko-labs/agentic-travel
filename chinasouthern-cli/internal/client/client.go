package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.csair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"China Southern",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "chinasouthern"), Brand: brand}
}
