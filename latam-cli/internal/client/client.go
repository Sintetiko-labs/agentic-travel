package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.latam.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"LATAM Airlines",
		"LATAM Brasil",
		"LATAM Chile",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "latam"), Brand: brand}
}
