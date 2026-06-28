package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.finnair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Finnair",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "finnair"), Brand: brand}
}
