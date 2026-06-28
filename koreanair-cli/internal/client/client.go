package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.koreanair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Korean Air",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "koreanair"), Brand: brand}
}
