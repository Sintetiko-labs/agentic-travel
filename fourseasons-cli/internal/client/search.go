package client

import "fmt"

// Search runs hotel search (TODO: implement for Four Seasons).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	_ = c
	return nil, fmt.Errorf("search not yet implemented for Four Seasons — see README and internal/client/search.go TODO")
}
