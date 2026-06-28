package client

import "fmt"

// Read returns flight or fare detail (stub).
func (c *Client) Read(idOrURL string) (*FlightView, error) {
	return nil, fmt.Errorf("read not yet implemented for World2Fly (id=%q)", idOrURL)
}
