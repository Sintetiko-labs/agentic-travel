package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.ihg.com"

// Client talks to IHG public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"IHG Hotels & Resorts",
		"InterContinental",
		"Kimpton",
		"Crowne Plaza",
		"Holiday Inn",
		"Holiday Inn Express",
		"Hotel Indigo",
		"Six Senses",
		"Vignette Collection",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "ihg"), Brand: brand}
}
