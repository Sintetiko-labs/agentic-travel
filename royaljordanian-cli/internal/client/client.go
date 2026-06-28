package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.rj.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Royal Jordanian",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "royaljordanian"), Brand: brand}
}
