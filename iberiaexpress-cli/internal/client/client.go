package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.iberiaexpress.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Iberia Express",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "iberiaexpress"), Brand: brand}
}
