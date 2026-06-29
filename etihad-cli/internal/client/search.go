package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	origin = strings.ToUpper(strings.TrimSpace(origin))
	dest = strings.ToUpper(strings.TrimSpace(dest))
	depart, err := parseYMD(depart)
	if err != nil {
		return nil, err
	}
	ret = strings.TrimSpace(ret)
	if ret != "" {
		if ret, err = parseYMD(ret); err != nil {
			return nil, err
		}
	}
	ref := c.BaseURL + "/es-es/"
	if _, err := c.FetchHTML(ref); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("etihad"))
		}
	}
	q := url.Values{}
	q.Set("origin", origin)
	q.Set("destination", dest)
	q.Set("departureDate", depart)
	q.Set("adults", "1")
	if ret != "" {
		q.Set("returnDate", ret)
	}
	body, status, err := c.apiGET(c.BaseURL+"/edge/deeplink/calendar?"+q.Encode(), ref)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	if err := checkAPI(status, body, "etihad"); err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	brand := c.Brand
	if brand == "" {
		brand = "Etihad Airways"
	}
	flights := flightsFromJSON(body, origin, dest, depart, brand)
	if flights == nil {
		flights = []FlightHit{}
	}
	return paginateFlightResult(origin, dest, depart, ret, page, pageSize, flights, brand, "deeplink/calendar"), nil
}
