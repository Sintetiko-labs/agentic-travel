package client

import "fmt"

// Search runs hotel search (TODO: implement for Relais & Châteaux).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	_ = c
	return nil, fmt.Errorf("search not yet implemented for Relais & Châteaux — see README and internal/client/search.go TODO")
}
