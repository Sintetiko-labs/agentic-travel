package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.avianca.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Avianca",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "avianca"), Brand: brand}
}
