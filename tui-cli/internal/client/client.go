package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.tui.co.uk"

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{
		"TUI Airways",
		"TUI fly",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "tui"), Brand: brand}
}
