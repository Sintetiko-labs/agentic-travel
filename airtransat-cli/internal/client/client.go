package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.airtransat.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air Transat",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "airtransat"), Brand: brand}
}
