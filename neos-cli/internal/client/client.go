package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.neosair.it"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Neos",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "neos"), Brand: brand}
}
