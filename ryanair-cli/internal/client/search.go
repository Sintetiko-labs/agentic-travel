package client

import (
	"fmt"
	"strings"
)

// Search runs flight search via booking availability API with farfnd fallback.
func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	origin = strings.ToUpper(strings.TrimSpace(origin))
	dest = strings.ToUpper(strings.TrimSpace(dest))
	depart = strings.TrimSpace(depart)
	ret = strings.TrimSpace(ret)
	if origin == "" || dest == "" || depart == "" {
		return nil, fmt.Errorf("origin, destination and depart date are required")
	}

	if res, err := c.searchAvailability(origin, dest, depart, ret); err == nil && res != nil && len(res.Flights) > 0 {
		res.Page = page
		res.PageSize = pageSize
		if c.Brand != "" {
			res.Brand = c.Brand
		}
		return paginateFlights(res, page, pageSize), nil
	}

	res, err := c.searchFarfnd(origin, dest, depart, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w — try RYANAIR_COOKIE or `ryanair session chrome` for booking API", origin, dest, err)
	}
	if c.Brand != "" {
		res.Brand = c.Brand
	}
	return res, nil
}

func paginateFlights(res *FlightSearchResult, page, pageSize int) *FlightSearchResult {
	total := len(res.Flights)
	start := (page - 1) * pageSize
	if start >= total {
		res.Flights = []FlightHit{}
	} else {
		end := start + pageSize
		if end > total {
			end = total
		}
		res.Flights = res.Flights[start:end]
	}
	res.Total = total
	res.Page = page
	res.PageSize = pageSize
	res.HasNext = total > page*pageSize
	return res
}
