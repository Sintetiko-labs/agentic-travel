package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.britishairways.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"British Airways",
		"BA CityFlyer",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "britishairways"), Brand: brand}
}
