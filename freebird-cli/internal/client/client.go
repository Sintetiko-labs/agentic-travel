package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.freebirdairlines.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Freebird Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "freebird"), Brand: brand}
}
