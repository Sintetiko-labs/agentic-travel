package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.vietnamairlines.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Vietnam Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "vietnamairlines"), Brand: brand}
}
