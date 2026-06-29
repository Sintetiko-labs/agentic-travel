package client

import "strings"

type brandRoute struct {
	name    string
	carrier string
}

func (c *Client) resolveBrand() brandRoute {
	brand := strings.TrimSpace(c.Brand)
	switch {
	case strings.EqualFold(brand, "Iberia Express"):
		return brandRoute{name: "Iberia Express", carrier: "I2"}
	case strings.EqualFold(brand, "Air Nostrum"):
		return brandRoute{name: "Air Nostrum", carrier: "YW"}
	default:
		return brandRoute{name: "Iberia", carrier: "IB"}
	}
}
