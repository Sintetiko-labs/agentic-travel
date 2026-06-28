package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.iberojet.es"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Iberojet",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "iberojet"), Brand: brand}
}
