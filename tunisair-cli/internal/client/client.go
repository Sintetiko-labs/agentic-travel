package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.tunisair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Tunisair",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "tunisair"), Brand: brand}
}
