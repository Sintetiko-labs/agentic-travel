package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.csa.cz"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Czech Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "czechairlines"), Brand: brand}
}
