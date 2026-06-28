package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

const (
	channelHomePath = "/api/channel-home/v1"
	bookingFlow     = "BOOKING"
)

// Search queries Air Europa via dapi channel-home (Amadeus BOOKING redirect + flightinfo fallback).
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

	if _, err := c.FetchHTML(BaseURL + "/es/es/"); err != nil {
		return nil, fmt.Errorf("bootstrap session: %w", err)
	}

	if res, err := c.searchBookingRedirect(origin, dest, depart, ret, page, pageSize); err == nil && res != nil {
		if c.Brand != "" {
			res.Brand = c.Brand
		}
		return res, nil
	} else if err != nil && strings.Contains(err.Error(), "akamai blocked") {
		return nil, err
	}

	res, err := c.searchFlightInfo(origin, dest, depart, ret, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w — try aireuropa session chrome --wait", origin, dest, err)
	}
	if c.Brand != "" {
		res.Brand = c.Brand
	}
	return res, nil
}

func (c *Client) searchBookingRedirect(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	payload := buildAmadeusBookingPayload(origin, dest, depart, ret)
	body, status, err := c.postDAPI(channelHomePath+"/redirect/flow/"+bookingFlow+"/urldata", payload, true)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		if akamai.IsWAFBlocked(status, string(body)) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("aireuropa"))
		}
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("akamai blocked or stale session (HTTP 404) — %s", akamai.NeedsSessionHint("aireuropa"))
		}
		return nil, fmt.Errorf("booking redirect HTTP %d: %s", status, tkbase.Truncate(string(body), 200))
	}
	if !akamai.LooksLikeJSON(string(body)) {
		return nil, fmt.Errorf("booking redirect: non-JSON response")
	}
	var resp amadeusRedirectResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	bookingURL := resp.bookingURL()
	if bookingURL == "" {
		return nil, fmt.Errorf("booking redirect: empty URL in response")
	}
	return &FlightSearchResult{
		Query:    fmt.Sprintf("%s-%s %s", origin, dest, depart),
		Origin:   origin,
		Dest:     dest,
		Depart:   depart,
		Return:   ret,
		Total:    1,
		Page:     page,
		PageSize: pageSize,
		Flights: []FlightHit{{
			ID:          fmt.Sprintf("%s-%s-%s", origin, dest, depart),
			Airline:     "Air Europa",
			Origin:      origin,
			Destination: dest,
			BookingURL:  bookingURL,
			Currency:    "EUR",
		}},
		Brand:  c.Brand,
		Source: "dapi-booking-redirect",
	}, nil
}

func (c *Client) searchFlightInfo(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	path := fmt.Sprintf("/api/channel/v2/flightinfo/airportDeparture/%s/airportArrival/%s", origin, dest)
	body, status, err := c.getDAPI(path)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		if akamai.IsWAFBlocked(status, string(body)) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("aireuropa"))
		}
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("akamai blocked or stale session (HTTP 404) — %s", akamai.NeedsSessionHint("aireuropa"))
		}
		return nil, fmt.Errorf("flightinfo HTTP %d: %s", status, tkbase.Truncate(string(body), 200))
	}
	if !akamai.LooksLikeJSON(string(body)) {
		return nil, fmt.Errorf("flightinfo: non-JSON response")
	}
	var resp aireuropaFlightInfoResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("flightinfo decode: %w", err)
	}
	return resp.toResult(origin, dest, depart, ret, page, pageSize, c.Brand), nil
}

func buildAmadeusBookingPayload(origin, dest, depart, ret string) map[string]string {
	dep, err := time.Parse("2006-01-02", depart)
	depStr := depart
	if err == nil {
		depStr = strings.ToUpper(dep.Format("02Jan06"))
	}
	trip := "O"
	if ret != "" {
		trip = "R"
	}
	return map[string]string{
		"B_ANY_TIME_1":             "TRUE",
		"B_ANY_TIME_2":             "TRUE",
		"B_DATE_1":                 depStr,
		"B_LOCATION_1":             origin,
		"B_LOCATION_2":             dest,
		"E_LOCATION_1":             dest,
		"E_LOCATION_2":             origin,
		"COMMERCIAL_FARE_FAMILY_1": "DIGITAL0",
		"DATE_RANGE_QUALIFIER_1":   "C",
		"DATE_RANGE_QUALIFIER_2":   "C",
		"TRIP_TYPE":                trip,
		"TRAVELLER_TYPE_1":         "ADT",
		"ENC_TIME":                 time.Now().UTC().Format("20060102150405"),
		"EXTERNAL_ID_10":           "governmentDiscount=false",
		"MARKET_CODE":              "ES",
		"SO_SITE_OFFICE_ID":        "MADUX08AA",
		"EXTERNAL_ID_12":           "device=DESKTOP",
		"VULCANO_PARAMS":           "resident=false,largeFamily=false,promocode=false,corporateCode=false,miles=false,multiCity=false",
		"CONFIG_SHOW_RSD_PRICE":    "TRUE",
	}
}

func (c *Client) postDAPI(path string, payload any, amadeusHeaders bool) ([]byte, int, error) {
	c.Throttle()
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequest(http.MethodPost, APIBaseURL+path, strings.NewReader(string(b)))
	if err != nil {
		return nil, 0, err
	}
	c.setDAPIHeaders(req, amadeusHeaders)
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	return body, resp.StatusCode, nil
}

func (c *Client) getDAPI(path string) ([]byte, int, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, APIBaseURL+path, nil)
	if err != nil {
		return nil, 0, err
	}
	c.setDAPIHeaders(req, false)
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	return body, resp.StatusCode, nil
}

func (c *Client) setDAPIHeaders(req *http.Request, amadeusHeaders bool) {
	c.SetAPIHeaders(req)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("origin", BaseURL)
	req.Header.Set("referer", BaseURL+"/es/es/")
	if amadeusHeaders {
		req.Header.Set("browser", "Chrome")
		req.Header.Set("browser-version", "131.0.0.0")
	}
}

type amadeusRedirectResponse struct {
	Data struct {
		URL string `json:"url"`
	} `json:"data"`
}

func (r *amadeusRedirectResponse) bookingURL() string {
	return r.Data.URL
}

type aireuropaFlightInfoResponse struct {
	Data []struct {
		FlightNumber      string `json:"flightNumber"`
		DepartureAirport  string `json:"departureAirportIataCode"`
		ArrivalAirport    string `json:"arrivalAirportIataCode"`
		DepartureDateTime string `json:"departureDateTime"`
		ArrivalDateTime   string `json:"arrivalDateTime"`
	} `json:"data"`
}

func (r *aireuropaFlightInfoResponse) toResult(origin, dest, depart, ret string, page, pageSize int, brand string) *FlightSearchResult {
	flights := make([]FlightHit, 0, len(r.Data))
	for _, f := range r.Data {
		dep := f.DepartureDateTime
		if i := strings.Index(dep, "T"); i > 0 {
			dep = dep[i+1:]
			if len(dep) >= 5 {
				dep = dep[:5]
			}
		}
		arr := f.ArrivalDateTime
		if i := strings.Index(arr, "T"); i > 0 {
			arr = arr[i+1:]
			if len(arr) >= 5 {
				arr = arr[:5]
			}
		}
		fn := f.FlightNumber
		if fn == "" {
			fn = "UX"
		}
		flights = append(flights, FlightHit{
			ID: fn + "-" + depart, Airline: "Air Europa", FlightNumber: fn,
			Origin: tkbase.FirstNonEmpty(f.DepartureAirport, origin),
			Destination: tkbase.FirstNonEmpty(f.ArrivalAirport, dest),
			Depart: dep, Arrive: arr,
			BookingURL: BaseURL + "/es/es/", Currency: "EUR",
		})
	}
	total := len(flights)
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}
	return &FlightSearchResult{
		Query: fmt.Sprintf("%s-%s %s", origin, dest, depart),
		Origin: origin, Dest: dest, Depart: depart, Return: ret,
		Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Flights: flights[start:end],
		Brand: brand, Source: "dapi-flightinfo",
	}
}
