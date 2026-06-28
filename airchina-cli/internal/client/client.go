package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.airchina.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air China",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "airchina"), Brand: brand}
}
