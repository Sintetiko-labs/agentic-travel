package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.cathaypacific.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Cathay Pacific",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "cathaypacific"), Brand: brand}
}
