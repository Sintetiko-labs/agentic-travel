package client

import (
	"encoding/json"
	"fmt"
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
	depart = strings.TrimSpace(depart)
	ret = strings.TrimSpace(ret)
	ref := c.BaseURL + "/es-int/"
	if _, err := c.FetchHTML(ref); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("turkish"))
		}
	}
	payload, _ := json.Marshal(map[string]any{
		"origin": origin, "destination": dest, "departureDate": depart, "returnDate": ret,
		"passengerCounts": map[string]int{"ADT": 1}, "cabin": "ECONOMY",
	})
	body, status, err := c.apiPOST(c.BaseURL+"/api/v1/availability", ref, payload)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	if err := checkAPI(status, body, "turkish"); err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	brand := c.Brand
	if brand == "" {
		brand = "Turkish Airlines"
	}
	flights := flightsFromJSON(body, origin, dest, depart, brand)
	if flights == nil {
		flights = []FlightHit{}
	}
	return paginateFlightResult(origin, dest, depart, ret, page, pageSize, flights, brand, "availability"), nil
}
