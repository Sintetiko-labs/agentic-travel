package client

import tkbase "github.com/fbelchi/travelkit/base"

const (
	BaseURL      = "https://www.plusultra.com"
	PricingAPI   = BaseURL + "/pricing/api/v1"
	PricingKey   = "d9eaa63c9008987381860a36e0d8c2aa2c6a936b41bf35e42bbe11e97bd452ea"
	CustomerGUID = "0f9ef31c-e69b-43c0-89c7-b2a7a0356d67"
)

type Client struct {
	*tkbase.Client
	Brand string
}

var Brands = []string{"Plus Ultra Líneas Aéreas"}

func New(brand string) *Client {
	return &Client{Client: tkbase.New(BaseURL, "plusultra"), Brand: brand}
}
