package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.wamosair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Wamos Air",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "wamos"), Brand: brand}
}
