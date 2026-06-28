package client

import "strings"

func (c *Client) isEurowingsBrand() bool {
	return strings.EqualFold(strings.TrimSpace(c.Brand), "Eurowings")
}

func (c *Client) defaultLHBrand() string {
	if b := strings.TrimSpace(c.Brand); b != "" {
		return b
	}
	return "Lufthansa"
}
