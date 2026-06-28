package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.airfrance.com"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"Air France",
		"KLM",
		"Transavia",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "airfranceklm"), Brand: brand}
}
