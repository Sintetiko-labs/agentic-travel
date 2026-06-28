package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.albastar.es"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Albastar",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "albastar"), Brand: brand}
}
