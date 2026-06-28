package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.flyplay.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"PLAY Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "play"), Brand: brand}
}
