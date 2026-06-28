package client

import tkbase "github.com/fbelchi/travelkit/base"

const BaseURL = "https://www.w2fly.es"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{"World2Fly"}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "world2fly"), Brand: brand}
}
