package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.iberia.es"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Iberia",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "iberia"), Brand: brand}
}
