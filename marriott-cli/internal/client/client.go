package client

import (
	tkbase "github.com/fbelchi/travelkit/base"
)

const BaseURL = "https://www.marriott.com"

// Client talks to Marriott public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
		"Marriott",
		"Marriott Hotels",
		"JW Marriott",
		"The Ritz-Carlton",
		"St. Regis",
		"W Hotels",
		"Edition",
		"Luxury Collection",
		"Westin",
		"Sheraton",
		"Le Méridien",
		"Renaissance Hotels",
		"Autograph Collection",
		"Tribute Portfolio",
		"AC Hotels",
		"AC Hotels by Marriott",
		"Aloft",
		"Moxy",
		"Courtyard by Marriott",
		"Residence Inn",
}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "marriott"), Brand: brand}
}
