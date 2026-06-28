package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.etihad.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Etihad Airways",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "etihad"), Brand: brand}
}
