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

const iberiaAvail = "https://www.iberia.com/api/availability/v1/flights"

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

	if _, err := c.FetchHTML("https://www.iberia.com/es/"); err != nil {
		return nil, fmt.Errorf("bootstrap: %w", err)
	}
	q := fmt.Sprintf("?market=ES&language=es&origin=%s&destination=%s&departureDate=%s&adults=1&operatingCarrier=IB", origin, dest, depart)
	c.Throttle()
	req, _ := http.NewRequest(http.MethodGet, iberiaAvail+q, nil)
	c.SetAPIHeaders(req)
	req.Header.Set("accept", "application/json")
	req.Header.Set("origin", "https://www.iberia.com")
	req.Header.Set("referer", "https://www.iberia.com/es/")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if akamai.IsWAFBlocked(resp.StatusCode, string(raw)) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("iberia"))
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("search %s→%s: HTTP %d", origin, dest, resp.StatusCode)
	}
	var av struct {
		Offers, Flights []struct {
			ID, FlightNumber, Origin, Destination, Departure, Arrival, Currency string
			Price                                                                 float64
		}
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
		flights = append(flights, FlightHit{ID: f.ID, Airline: "Iberia", FlightNumber: f.FlightNumber, Origin: f.Origin, Destination: f.Destination, Depart: f.Departure, Arrive: f.Arrival, Price: fmt.Sprintf("%.2f", f.Price), Currency: tkbase.FirstNonEmpty(f.Currency, "EUR"), BookingURL: "https://www.iberia.com/es/vuelos/"})
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
	return &FlightSearchResult{Query: origin + "-" + dest + " " + depart, Origin: origin, Dest: dest, Depart: depart, Return: ret, Total: total, Page: page, PageSize: pageSize, Flights: pf, Brand: c.Brand, Source: "iberia-availability-yw"}, nil
}
