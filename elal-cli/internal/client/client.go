package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.elal.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"El Al",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "elal"), Brand: brand}
}
