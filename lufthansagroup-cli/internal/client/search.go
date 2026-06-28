package client

import "fmt"

// Search runs flight search (TODO: implement for Lufthansa Group).
func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	_ = c
	return nil, fmt.Errorf("search not yet implemented for Lufthansa Group — see README and internal/client/search.go TODO")
}
