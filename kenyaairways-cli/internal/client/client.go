package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.kenya-airways.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Kenya Airways",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "kenyaairways"), Brand: brand}
}
