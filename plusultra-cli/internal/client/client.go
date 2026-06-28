package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.plusultra.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Plus Ultra Líneas Aéreas",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "plusultra"), Brand: brand}
}
