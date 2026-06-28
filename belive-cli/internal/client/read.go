package client

import "fmt"

// Read returns hotel detail (stub).
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	return nil, fmt.Errorf("read not yet implemented for Be Live (id=%q)", idOrURL)
}
