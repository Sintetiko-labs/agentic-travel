package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.airnostrum.es"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air Nostrum",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "airnostrum"), Brand: brand}
}
