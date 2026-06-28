package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.nouvelair.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Nouvelair",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "nouvelair"), Brand: brand}
}
