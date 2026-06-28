package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.croatiaairlines.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Croatia Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "croatiaairlines"), Brand: brand}
}
