package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.united.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"United Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "united"), Brand: brand}
}
