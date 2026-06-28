package client

import "fmt"

func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	_ = c
	return nil, fmt.Errorf("swiftair has no public scheduled passenger search API (cargo/ACMI only)")
}
