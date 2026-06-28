package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	if page < 1 { page = 1 }
	if pageSize < 1 { pageSize = 24 }
	origin = strings.ToUpper(strings.TrimSpace(origin))
	dest = strings.ToUpper(strings.TrimSpace(dest))
	depart, err := parseYMD(depart)
	if err != nil { return nil, err }
	ret = strings.TrimSpace(ret)
	if ret != "" { if ret, err = parseYMD(ret); err != nil { return nil, err } }
	ref := c.BaseURL + "/es/spanish/"
	if _, err := c.FetchHTML(ref); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("emirates"))
		}
	}
	segments := []map[string]any{{"origin": origin, "destination": dest, "departureDate": depart}}
	if ret != "" { segments = append(segments, map[string]any{"origin": dest, "destination": origin, "departureDate": ret}) }
	payload, _ := json.Marshal(map[string]any{"segments": segments, "passengers": []map[string]any{{"type": "ADT", "count": 1}}, "bookingType": "REVENUE", "cabinClass": "Y"})
	body, status, err := c.apiPOST(c.BaseURL+"/service/search/search-results", ref, payload)
	if err != nil { return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err) }
	if err := checkAPI(status, body, "emirates"); err != nil { return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err) }
	brand := c.Brand
	if brand == "" { brand = "Emirates" }
	flights := flightsFromJSON(body, origin, dest, depart, brand)
	if flights == nil { flights = []FlightHit{} }
	for i := range flights {
		if flights[i].BookingURL == "" {
			q := url.Values{}; q.Set("from", origin); q.Set("to", dest); q.Set("depart", depart)
			flights[i].BookingURL = c.BaseURL + "/es/spanish/book?" + q.Encode()
		}
	}
	return paginateFlightResult(origin, dest, depart, ret, page, pageSize, flights, brand, "search-results"), nil
}
