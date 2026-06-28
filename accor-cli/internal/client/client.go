package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://all.accor.com"

// Client talks to Accor public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Ibis",
		"Ibis Budget",
		"Ibis Styles",
		"Novotel",
		"Mercure",
		"Pullman",
		"Sofitel",
		"MGallery",
		"Fairmont",
		"Raffles",
		"Accor",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "accor"), Brand: brand}
}
