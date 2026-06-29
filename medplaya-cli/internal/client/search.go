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
	return tkhotel.SearchSpanishHTML(c.FetchHTML, tkhotel.SpanishSearchOpts{
		Query: query, Brand: c.Brand, BaseURL: c.BaseURL, Source: "html-links",
		DefaultBrand: "MedPlaya", Paths: tkhotel.EsHotelPaths(query),
		Parse: parse.HotelsFromMedplayaLinks, Page: page, PageSize: pageSize,
	})
}
