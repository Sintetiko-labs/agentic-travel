package client

import "fmt"

// Search runs hotel search (TODO: implement for Derby Hotels).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	_ = c
	return nil, fmt.Errorf("search not yet implemented for Derby Hotels — see README and internal/client/search.go TODO")
}
