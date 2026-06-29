package client

import (
	"strings"

	tkhotel "github.com/fbelchi/travelkit/hotel"
	"github.com/fbelchi/travelkit/parse"
)

func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	brand := c.Brand
	if brand == "" {
		brand = "Bahia Principe"
	}
	return tkhotel.SearchSpanishHTML(c.FetchHTML, tkhotel.SpanishSearchOpts{
		Query: query, Brand: brand, BaseURL: c.BaseURL, Source: "html-links",
		DefaultBrand: "Bahia Principe", Paths: tkhotel.EsHotelPaths(query),
		Parse: parse.HotelsFromBahiaLinks, Page: page, PageSize: pageSize,
	})
}
