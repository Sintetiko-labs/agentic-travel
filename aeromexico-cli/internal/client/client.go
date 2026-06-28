package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.aeromexico.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Aeroméxico",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "aeromexico"), Brand: brand}
}
