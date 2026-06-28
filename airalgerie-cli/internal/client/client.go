package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.airalgerie.dz"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air Algerie",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "airalgerie"), Brand: brand}
}
