package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.gulfair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Gulf Air",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "gulfair"), Brand: brand}
}
