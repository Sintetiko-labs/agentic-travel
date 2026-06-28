package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.tochostels.com"

// Client talks to TOC Hostels public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"TOC Hostels",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "toc"), Brand: brand}
}
