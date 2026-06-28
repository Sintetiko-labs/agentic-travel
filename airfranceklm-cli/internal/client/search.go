package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
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
	switch c.resolveAFKLMBrand() {
	case "KLM":
		return c.searchKLM(origin, dest, depart, ret, page, pageSize, "KLM")
	case "Transavia":
		return c.searchTransavia(origin, dest, depart, ret, page, pageSize)
	default:
		return c.searchAirFrance(origin, dest, depart, ret, page, pageSize)
	}
}


func (c *Client) searchAirFrance(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	ref := c.BaseURL + "/es/"
	if _, err := c.FetchHTML(ref); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("airfranceklm"))
		}
	}
	payload, _ := json.Marshal(map[string]any{
		"commercialCabins": []string{"ECONOMY"},
		"itineraries":      []map[string]any{{"origin": origin, "destination": dest, "departureDate": depart}},
		"passengers":       []map[string]any{{"id": 1, "type": "ADT"}},
	})
	body, status, err := c.apiPOST(c.BaseURL+"/api/v1/search/air-bounds", ref, payload)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	if err := checkAPI(status, body, "airfranceklm"); err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	brand := "Air France"
	flights := flightsFromJSON(body, origin, dest, depart, brand)
	if flights == nil {
		flights = []FlightHit{}
	}
	return paginateFlightResult(origin, dest, depart, ret, page, pageSize, flights, brand, "air-bounds"), nil
}

func (c *Client) searchKLM(origin, dest, depart, ret string, page, pageSize int, brand string) (*FlightSearchResult, error) {
	ref := klmBase + "/es/"
	if _, err := c.FetchHTML(ref); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("airfranceklm"))
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
	body, status, err := c.apiGET(klmBase+"/api/calendar?"+q.Encode(), ref)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	if err := checkAPI(status, body, "airfranceklm"); err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	flights := flightsFromJSON(body, origin, dest, depart, brand)
	if flights == nil {
		flights = []FlightHit{}
	}
	return paginateFlightResult(origin, dest, depart, ret, page, pageSize, flights, brand, "klm/calendar"), nil
}

func (c *Client) searchTransavia(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	ref := transaviaBase + "/es-es/"
	if _, err := c.FetchHTML(ref); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("airfranceklm"))
		}
	}
	q := url.Values{}
	q.Set("from", origin)
	q.Set("to", dest)
	q.Set("departureDate", depart)
	q.Set("adults", "1")
	if ret != "" {
		q.Set("returnDate", ret)
	}
	body, status, err := c.apiGET(transaviaBase+"/api/offers?"+q.Encode(), ref)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	if err := checkAPI(status, body, "airfranceklm"); err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	brand := "Transavia"
	flights := flightsFromJSON(body, origin, dest, depart, brand)
	if flights == nil {
		flights = []FlightHit{}
	}
	return paginateFlightResult(origin, dest, depart, ret, page, pageSize, flights, brand, "transavia/offers"), nil
}
