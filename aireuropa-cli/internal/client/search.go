package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search queries Air Europa flight search API.
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

	payload := map[string]any{
		"origin": origin, "destination": dest,
		"departureDate": depart, "returnDate": ret,
		"adults": 1, "language": "es", "market": "ES",
	}
	var resp aireuropaSearchResponse
	if err := c.PostJSON(c.BaseURL+"/ae/api/v1/flights/search", payload, &resp); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("aireuropa"))
		}
		return nil, fmt.Errorf("search %s→%s: %w — try AIREUROPA_COOKIE", origin, dest, err)
	}
	return resp.toResult(origin, dest, depart, ret, page, pageSize, c.Brand), nil
}

type aireuropaSearchResponse struct {
	Flights []struct {
		ID           string  `json:"id"`
		FlightNumber string  `json:"flightNumber"`
		Origin       string  `json:"origin"`
		Destination  string  `json:"destination"`
		Departure    string  `json:"departure"`
		Arrival      string  `json:"arrival"`
		Duration     string  `json:"duration"`
		Price        float64 `json:"price"`
		Currency     string  `json:"currency"`
		Stops        int     `json:"stops"`
	} `json:"flights"`
	Total int `json:"total"`
}

func (r *aireuropaSearchResponse) toResult(origin, dest, depart, ret string, page, pageSize int, brand string) *FlightSearchResult {
	flights := make([]FlightHit, 0, len(r.Flights))
	for _, f := range r.Flights {
		flights = append(flights, FlightHit{
			ID: f.ID, Airline: "Air Europa", FlightNumber: f.FlightNumber,
			Origin: f.Origin, Destination: f.Destination,
			Depart: f.Departure, Arrive: f.Arrival, Duration: f.Duration,
			Stops: f.Stops, Price: fmt.Sprintf("%.2f", f.Price), Currency: f.Currency,
			BookingURL: "https://www.aireuropa.com/es/es/",
		})
	}
	total := r.Total
	if total == 0 {
		total = len(flights)
	}
	return &FlightSearchResult{
		Query: fmt.Sprintf("%s-%s %s", origin, dest, depart),
		Origin: origin, Dest: dest, Depart: depart, Return: ret,
		Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Flights: flights,
		Brand: brand, Source: "api",
	}
}
