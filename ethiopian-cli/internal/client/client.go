package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.ethiopianairlines.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Ethiopian Airlines",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "ethiopian"), Brand: brand}
}
