package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.vueling.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Vueling",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "vueling"), Brand: brand}
}
