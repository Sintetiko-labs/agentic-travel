package client

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

const wizzCulture = "es-es"

var wizzVersionRE = regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)

type wizzSearchRequest struct {
	IsFlightChange bool              `json:"isFlightChange"`
	IsReturnFlight bool              `json:"isReturnFlight"`
	AdultCount     int               `json:"adultCount"`
	ChildCount     int               `json:"childCount"`
	InfantCount    int               `json:"infantCount"`
	FlightList     []wizzFlightSlice `json:"flightList"`
}

type wizzFlightSlice struct {
	DepartureStation string `json:"departureStation"`
	ArrivalStation   string `json:"arrivalStation"`
	DepartureDate    string `json:"departureDate"`
}

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

	ref := c.BaseURL + "/" + wizzCulture
	if _, err := c.FetchHTML(ref); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("wizzair"))
		}
		return nil, fmt.Errorf("bootstrap: %w", err)
	}

	version, err := c.fetchAPIVersion(ref)
	if err != nil {
		return nil, err
	}

	reqBody := wizzSearchRequest{
		IsFlightChange: false,
		IsReturnFlight: ret != "",
		AdultCount:     1,
		FlightList: []wizzFlightSlice{{
			DepartureStation: origin,
			ArrivalStation:   dest,
			DepartureDate:    wizzDepartDate(depart),
		}},
	}
	if ret != "" {
		reqBody.FlightList = append(reqBody.FlightList, wizzFlightSlice{
			DepartureStation: dest,
			ArrivalStation:   origin,
			DepartureDate:    wizzDepartDate(ret),
		})
	}
	retSeg := "null"
	if ret != "" {
		retSeg = ret
	}
	payload, _ := json.Marshal(reqBody)
	apiURL := fmt.Sprintf("https://be.wizzair.com/%s/Api/search/search", version)
	bookRef := fmt.Sprintf("%s/#/booking/select-flight/%s/%s/%s/%s/1/0/0/null",
		c.BaseURL, origin, dest, depart, retSeg)
	body, status, err := c.wizzPOST(apiURL, bookRef, payload)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	if err := checkAPI(status, body, "wizzair"); err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	brand := c.Brand
	if brand == "" {
		brand = "Wizz Air"
	}
	flights := flightsFromWizzJSON(body, origin, dest, depart, brand)
	if flights == nil {
		flights = []FlightHit{}
	}
	return paginateFlightResult(origin, dest, depart, ret, page, pageSize, flights, brand, "Api/search/search"), nil
}

func wizzDepartDate(depart string) string {
	if strings.Contains(depart, "T") {
		return depart
	}
	return depart + "T00:00:00"
}

func (c *Client) fetchAPIVersion(referer string) (string, error) {
	body, status, err := c.apiGET(c.BaseURL+"/buildnumberjson", referer)
	if err != nil {
		return "", err
	}
	if err := checkAPI(status, body, "wizzair"); err != nil {
		return "", fmt.Errorf("buildnumberjson: %w", err)
	}
	var meta map[string]any
	if err := json.Unmarshal(body, &meta); err != nil {
		if m := wizzVersionRE.FindString(string(body)); m != "" {
			return m, nil
		}
		return "", fmt.Errorf("buildnumberjson decode: %w", err)
	}
	for _, k := range []string{"buildId", "buildNumber", "version", "BuildId", "BuildNumber"} {
		if v, ok := meta[k]; ok {
			if s, ok := v.(string); ok && s != "" {
				if m := wizzVersionRE.FindString(s); m != "" {
					return m, nil
				}
				return s, nil
			}
		}
	}
	if m := wizzVersionRE.FindString(string(body)); m != "" {
		return m, nil
	}
	return "", fmt.Errorf("buildnumberjson: no version in response")
}
