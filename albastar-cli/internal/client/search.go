package client

import "fmt"

func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	_ = c
	return nil, fmt.Errorf("albastar has no public scheduled passenger search API (charter only)")
}
