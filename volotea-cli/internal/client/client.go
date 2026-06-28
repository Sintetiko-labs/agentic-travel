package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.volotea.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Volotea",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "volotea"), Brand: brand}
}
