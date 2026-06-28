package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.condor.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Condor",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "condor"), Brand: brand}
}
