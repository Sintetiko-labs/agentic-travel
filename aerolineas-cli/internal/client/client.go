package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.aerolineas.com.ar"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Aerolíneas Argentinas",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "aerolineas"), Brand: brand}
}
