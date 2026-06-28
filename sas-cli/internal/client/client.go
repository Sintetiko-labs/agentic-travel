package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.flysas.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"SAS",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "sas"), Brand: brand}
}
