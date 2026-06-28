package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

const levelAPI = "https://lv-uihomeweb-webapp-pro.azurewebsites.net/api/v1/search/flights"

func (c *Client) Search(origin, dest, depart, ret string, page, pageSize int) (*FlightSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	origin = strings.ToUpper(strings.TrimSpace(origin))
	dest = strings.ToUpper(strings.TrimSpace(dest))
	depart, ret = strings.TrimSpace(depart), strings.TrimSpace(ret)

	if _, err := c.FetchHTML(BaseURL + "/es/"); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("level"))
		}
	}
	payload := map[string]any{"origin": origin, "destination": dest, "departureDate": depart, "adults": 1, "culture": "es-ES", "currency": "EUR", "tripType": "OneWay"}
	if ret != "" {
		payload["returnDate"] = ret
		payload["tripType"] = "RoundTrip"
	}
	b, _ := json.Marshal(payload)
	c.Throttle()
	req, _ := http.NewRequest(http.MethodPost, levelAPI, strings.NewReader(string(b)))
	c.SetAPIHeaders(req)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("origin", BaseURL)
	req.Header.Set("referer", BaseURL+"/es/")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if resp.StatusCode == 403 || akamai.IsWAFBlocked(resp.StatusCode, string(raw)) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("level"))
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("search %s→%s: HTTP %d", origin, dest, resp.StatusCode)
	}
	var r struct {
		Flights []struct {
			FlightNumber, Origin, Destination, Departure, Arrival, Currency string
			Price                                                           float64
		} `json:"flights"`
	}
	_ = json.Unmarshal(raw, &r)
	flights := []FlightHit{}
	for _, f := range r.Flights {
		flights = append(flights, FlightHit{ID: f.FlightNumber, Airline: "Level", FlightNumber: f.FlightNumber, Origin: origin, Destination: dest, Depart: f.Departure, Arrive: f.Arrival, Price: fmt.Sprintf("%.2f", f.Price), Currency: "EUR", BookingURL: BaseURL + "/es/"})
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
	pf := flights[start:end]
	if pf == nil {
		pf = []FlightHit{}
	}
	return &FlightSearchResult{Query: origin + "-" + dest + " " + depart, Origin: origin, Dest: dest, Depart: depart, Return: ret, Total: total, Page: page, PageSize: pageSize, Flights: pf, Brand: c.Brand, Source: "uihomeweb/search"}, nil
}
