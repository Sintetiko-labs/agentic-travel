package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.aircanada.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air Canada",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "aircanada"), Brand: brand}
}
