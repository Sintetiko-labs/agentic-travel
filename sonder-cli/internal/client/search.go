package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
	tkhotel "github.com/fbelchi/travelkit/hotel"
	"github.com/fbelchi/travelkit/parse"
)

const sonderGraphURL = "https://graph.sonder.com"

func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, fmt.Errorf("destination required")
	}
	rows, err := c.sonderGraphQLSearch(query)
	if err != nil || len(rows) == 0 {
		html, ferr := c.FetchHTML(c.BaseURL + "/")
		if ferr != nil {
			if err != nil {
				return nil, err
			}
			if he, ok := ferr.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
				return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("sonder"))
			}
			return nil, fmt.Errorf("search %q: %w", query, ferr)
		}
		if akamai.IsDenied(403, html) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("sonder"))
		}
		rows = parse.HotelsFromSonderHome(html, c.BaseURL)
		rows = tkhotel.FilterHotelLD(rows, query)
		if len(rows) == 0 {
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("search %q: no hotels parsed", query)
		}
		b := c.Brand
		if b == "" {
			b = "Sonder"
		}
		return tkhotel.LDToResult(rows, query, page, pageSize, b, c.BaseURL, "homepage"), nil
	}
	rows = tkhotel.FilterHotelLD(rows, query)
	b := c.Brand
	if b == "" {
		b = "Sonder"
	}
	return tkhotel.LDToResult(rows, query, page, pageSize, b, c.BaseURL, "graphql"), nil
}

func (c *Client) sonderGraphQLSearch(query string) ([]parse.HotelLD, error) {
	payload := map[string]any{
		"operationName": "buildingSearch",
		"query": `query buildingSearch($filters: BuildingSearchFilters!) {
  buildingSearch(filters: $filters) {
    buildings { id name slug city country }
  }
}`,
		"variables": map[string]any{"filters": map[string]any{"query": strings.TrimSpace(query)}},
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	c.Throttle()
	req, err := http.NewRequest(http.MethodPost, sonderGraphURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	c.SetAPIHeaders(req)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", c.BaseURL)
	req.Header.Set("referer", c.BaseURL+"/")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if akamai.IsDenied(resp.StatusCode, string(body)) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("sonder"))
		}
		return nil, fmt.Errorf("graphql: HTTP %d", resp.StatusCode)
	}
	return parse.HotelsFromSonderGraphQL(body, c.BaseURL), nil
}
