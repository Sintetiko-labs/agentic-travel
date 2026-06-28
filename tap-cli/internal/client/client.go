package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.flytap.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"TAP Air Portugal",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "tap"), Brand: brand}
}
