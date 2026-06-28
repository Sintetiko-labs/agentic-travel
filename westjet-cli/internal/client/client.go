package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.westjet.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"WestJet",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "westjet"), Brand: brand}
}
