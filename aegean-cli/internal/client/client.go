package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.aegeanair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Aegean Airlines",
		"Olympic Air",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "aegean"), Brand: brand}
}
