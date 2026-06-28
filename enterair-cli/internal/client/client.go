package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.enterair.pl"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Enter Air",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "enterair"), Brand: brand}
}
