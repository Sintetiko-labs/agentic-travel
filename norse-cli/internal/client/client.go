package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.flynorse.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Norse Atlantic Airways",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "norse"), Brand: brand}
}
