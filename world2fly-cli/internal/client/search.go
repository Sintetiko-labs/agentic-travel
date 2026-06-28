package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	depart, ret = strings.TrimSpace(depart), strings.TrimSpace(ret)

	if _, err := c.FetchHTML(BaseURL + "/es-es/"); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && he.Status == 403 {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("world2fly"))
		}
	}
	q := fmt.Sprintf("/api/availability/v1/flights?market=ES&language=es&origin=%s&destination=%s&departureDate=%s&adults=1&operatingCarrier=2W", origin, dest, depart)
	if ret != "" {
		q += "&returnDate=" + url.QueryEscape(ret)
	}
	c.Throttle()
	req, _ := http.NewRequest(http.MethodGet, c.BaseURL+q, nil)
	c.SetAPIHeaders(req)
	req.Header.Set("accept", "application/json")
	req.Header.Set("origin", c.BaseURL)
	req.Header.Set("referer", c.BaseURL+"/es-es/")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if akamai.IsWAFBlocked(resp.StatusCode, string(raw)) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("world2fly"))
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("search %s→%s: HTTP %d", origin, dest, resp.StatusCode)
	}
	var av struct {
		Offers []flightOffer `json:"offers"`
		Flights []flightOffer `json:"flights"`
	}
	if err := json.Unmarshal(raw, &av); err != nil {
		return nil, err
	}
	items := av.Offers
	if len(items) == 0 {
		items = av.Flights
	}
	flights := make([]FlightHit, 0, len(items))
	for _, f := range items {
		flights = append(flights, FlightHit{ID: f.ID, Airline: "World2Fly", FlightNumber: f.FlightNumber, Origin: f.Origin, Destination: f.Destination, Depart: f.Departure, Arrive: f.Arrival, Price: fmt.Sprintf("%.2f", f.Price), Currency: tkbase.FirstNonEmpty(f.Currency, "EUR"), BookingURL: BaseURL + "/es-es/"})
	}
	return paginateW2(origin, dest, depart, ret, page, pageSize, flights, c.Brand), nil
}

type flightOffer struct {
	ID           string  `json:"id"`
	FlightNumber string  `json:"flightNumber"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	Departure    string  `json:"departure"`
	Arrival      string  `json:"arrival"`
	Price        float64 `json:"price"`
	Currency     string  `json:"currency"`
}

func paginateW2(origin, dest, depart, ret string, page, pageSize int, flights []FlightHit, brand string) *FlightSearchResult {
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
	return &FlightSearchResult{Query: origin + "-" + dest + " " + depart, Origin: origin, Dest: dest, Depart: depart, Return: ret, Total: total, Page: page, PageSize: pageSize, HasNext: total > page*pageSize, Flights: pf, Brand: brand, Source: "availability"}
}
