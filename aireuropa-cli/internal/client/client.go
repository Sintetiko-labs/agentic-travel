package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.aireuropa.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air Europa",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "aireuropa"), Brand: brand}
}
