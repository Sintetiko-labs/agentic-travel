package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	tkbase "github.com/fbelchi/travelkit/base"
)

const (
	voloteaAPIBase = "https://api.volotea.com/api/spa/voe/v1"
	voloteaAPIKey  = "ec92495bea0f4405abb71df117-spdot"
	voloteaBookURL = "https://book.volotea.com"
)

// Search queries Volotea booking SPA API (anonymous login + flights/search).
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

	token, err := c.voloteaLogin()
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: login: %w", origin, dest, err)
	}

	body, err := c.voloteaSearch(token, origin, dest, depart, ret)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}

	flights := parseVoloteaSearch(body, origin, dest, depart, ret)
	total := len(flights)
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}
	pageFlights := flights[start:end]
	if pageFlights == nil {
		pageFlights = []FlightHit{}
	}

	return &FlightSearchResult{
		Query:    fmt.Sprintf("%s-%s %s", origin, dest, depart),
		Origin:   origin,
		Dest:     dest,
		Depart:   depart,
		Return:   ret,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		HasNext:  total > page*pageSize,
		Flights:  pageFlights,
		Brand:    c.Brand,
		Source:   "flights/search",
	}, nil
}

func (c *Client) voloteaLogin() (string, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodPost, voloteaAPIBase+"/account/login",
		strings.NewReader(`{"isRemembered":true,"hash":"","username":"","password":""}`))
	if err != nil {
		return "", err
	}
	c.setVoloteaHeaders(req, "")
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", &tkbase.HTTPError{Status: resp.StatusCode, Body: tkbase.Truncate(string(raw), 300)}
	}
	var out struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return "", err
	}
	if out.Data.Token == "" {
		return "", fmt.Errorf("empty session token")
	}
	return out.Data.Token, nil
}

func (c *Client) voloteaSearch(token, origin, dest, depart, ret string) ([]byte, error) {
	payload := voloteaSearchPayload(origin, dest, depart, ret)
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	c.Throttle()
	req, err := http.NewRequest(http.MethodPost, voloteaAPIBase+"/flights/search", strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	c.setVoloteaHeaders(req, token)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &tkbase.HTTPError{Status: resp.StatusCode, Body: tkbase.Truncate(string(raw), 300)}
	}
	return raw, nil
}

func (c *Client) setVoloteaHeaders(req *http.Request, token string) {
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-api-key", voloteaAPIKey)
	req.Header.Set("auth", "true")
	req.Header.Set("x-client-session", "true")
	req.Header.Set("origin", voloteaBookURL)
	req.Header.Set("referer", voloteaBookURL+"/")
	c.ApplyCookie(req)
	if token != "" {
		req.Header.Set("x-session-token", token)
	}
}

func voloteaSearchPayload(origin, dest, depart, ret string) map[string]any {
	d, _ := time.Parse("2006-01-02", depart)
	begin := d.AddDate(0, 0, -9).Format("2006-01-02")
	end := d.AddDate(0, 0, 10).Format("2006-01-02")
	bt := "OneWay"
	if ret != "" {
		bt = "Return"
	}
	return map[string]any{
		"criteria": []map[string]any{{
			"beginDate": begin, "endDate": end, "selectedDate": depart,
			"origin": origin, "destination": dest,
		}},
		"codes": map[string]any{
			"currency": "EUR", "promotionCode": "", "bookingType": bt, "residentType": "NONE",
		},
		"passengers":         []map[string]any{{"type": "ADT", "count": 1}},
		"fareTypesToRequest": []string{"R"},
	}
}

type voloteaSearchResponse struct {
	Data struct {
		Trips []struct {
			JourneysAvailable []voloteaJourney `json:"journeysAvailable"`
		} `json:"trips"`
	} `json:"data"`
}

type voloteaJourney struct {
	HasFlights bool `json:"hasFlights"`
	Segments   []struct {
		Designator struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Departure   string `json:"departure"`
			Arrival     string `json:"arrival"`
		} `json:"designator"`
		Identifier struct {
			CarrierCode string `json:"carrierCode"`
			Identifier  string `json:"identifier"`
		} `json:"identifier"`
	} `json:"segments"`
	Fares []struct {
		PassengerFares []struct {
			FareAmount struct {
				Amount float64 `json:"amount"`
			} `json:"fareAmount"`
		} `json:"passengerFares"`
	} `json:"fares"`
}

func parseVoloteaSearch(raw []byte, origin, dest, depart, ret string) []FlightHit {
	var resp voloteaSearchResponse
	if err := jsonUnmarshal(raw, &resp); err != nil {
		return []FlightHit{}
	}
	return filterVoloteaJourneys(resp, origin, dest, depart, ret)
}

func filterVoloteaJourneys(resp voloteaSearchResponse, origin, dest, depart, ret string) []FlightHit {
	var out = make([]FlightHit, 0)
	for _, trip := range resp.Data.Trips {
		for _, j := range trip.JourneysAvailable {
			if !j.HasFlights || len(j.Segments) == 0 || len(j.Fares) == 0 {
				continue
			}
			seg := j.Segments[0]
			depDay := strings.Split(seg.Designator.Departure, "T")[0]
			if depDay != depart {
				continue
			}
			if !strings.EqualFold(seg.Designator.Origin, origin) || !strings.EqualFold(seg.Designator.Destination, dest) {
				continue
			}
			price := ""
			if len(j.Fares[0].PassengerFares) > 0 {
				price = fmt.Sprintf("%.2f", j.Fares[0].PassengerFares[0].FareAmount.Amount)
			}
			fn := voloteaFlightNumber(seg.Identifier.CarrierCode, seg.Identifier.Identifier)
			out = append(out, FlightHit{
				ID:           fmt.Sprintf("%s-%s-%s-%s", origin, dest, depart, fn),
				Airline:      "Volotea",
				FlightNumber: fn,
				Origin:       origin,
				Destination:  dest,
				Depart:       voloteaTime(seg.Designator.Departure),
				Arrive:       voloteaTime(seg.Designator.Arrival),
				Price:        price,
				Currency:     "EUR",
				BookingURL:   voloteaBookingURL(origin, dest, depart, ret),
			})
		}
	}
	return out
}

func voloteaFlightNumber(carrier, id string) string {
	carrier = strings.ToUpper(strings.TrimSpace(carrier))
	id = strings.TrimSpace(id)
	switch carrier {
	case "V7", "VOLOTEA":
		return "V7" + id
	default:
		if carrier != "" {
			return carrier + id
		}
		return "V7" + id
	}
}

func voloteaTime(iso string) string {
	if i := strings.Index(iso, "T"); i >= 0 && len(iso) >= i+6 {
		return iso[i+1 : i+6]
	}
	return iso
}

func voloteaBookingURL(origin, dest, depart, ret string) string {
	q := url.Values{}
	q.Set("origin", origin)
	q.Set("destination", dest)
	q.Set("departureDate", depart)
	if ret != "" {
		q.Set("returnDate", ret)
	}
	return voloteaBookURL + "/?" + q.Encode()
}
