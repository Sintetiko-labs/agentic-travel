package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.ita-airways.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"ITA Airways",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "ita"), Brand: brand}
}
