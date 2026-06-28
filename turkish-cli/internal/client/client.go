package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.turkishairlines.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Turkish Airlines",
		"Pegasus Airlines",
		"SunExpress",
		"AnadoluJet / AJet",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "turkish"), Brand: brand}
}
