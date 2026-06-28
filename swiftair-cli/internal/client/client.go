package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.swiftair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Swiftair",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "swiftair"), Brand: brand}
}
