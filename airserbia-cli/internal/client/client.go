package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.airserbia.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air Serbia",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "airserbia"), Brand: brand}
}
