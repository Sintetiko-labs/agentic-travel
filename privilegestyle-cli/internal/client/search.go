package client

import "fmt"

func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	_ = c
	return nil, fmt.Errorf("privilege style has no public scheduled passenger search API (charter/ACMI only)")
}
