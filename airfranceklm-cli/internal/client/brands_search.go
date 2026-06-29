package client

import "strings"

func (c *Client) resolveAFKLMBrand() string {
	brand := strings.TrimSpace(c.Brand)
	switch {
	case brand == "" || strings.EqualFold(brand, "KLM"):
		return "KLM"
	case strings.EqualFold(brand, "Transavia"):
		return "Transavia"
	default:
		return "Air France"
	}
}
