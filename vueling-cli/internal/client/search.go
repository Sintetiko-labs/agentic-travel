package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search queries Vueling flight search API (Skysales BIT endpoint).
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

	q := url.Values{}
	q.Set("origin", origin)
	q.Set("destination", dest)
	q.Set("departureDate", depart)
	q.Set("adults", "1")
	if ret != "" {
		q.Set("returnDate", ret)
	}
	path := "/bit/v2/flights/search?" + q.Encode()

	var resp vuelingSearchResponse
	err := c.searchBIT(c.BaseURL+path, &resp)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && he.Status == 404 {
			// www.vueling.com serves SPA shell; Skysales BIT lives on tickets host.
			err = c.searchBIT(skysalesBaseURL+path, &resp)
		}
	}
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("vueling"))
		}
		return nil, fmt.Errorf("search %s→%s: %w — run `vueling session chrome` (BIT v2 on tickets.vueling.com)", origin, dest, err)
	}
	return resp.toResult(origin, dest, depart, ret, page, pageSize, c.Brand, c.BaseURL), nil
}

const skysalesBaseURL = "https://tickets.vueling.com"

func (c *Client) searchBIT(fullURL string, out *vuelingSearchResponse) error {
	body, status, err := c.GetRaw(fullURL)
	if err != nil {
		return err
	}
	if status < 200 || status >= 300 {
		return &tkbase.HTTPError{Status: status, Body: tkbase.Truncate(string(body), 300)}
	}
	if err := jsonUnmarshal(body, out); err != nil {
		return err
	}
	return nil
}

type vuelingSearchResponse struct {
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
	} `json:"flights"`
	Total int `json:"total"`
}

func (r *vuelingSearchResponse) toResult(origin, dest, depart, ret string, page, pageSize int, brand, base string) *FlightSearchResult {
	flights := make([]FlightHit, 0, len(r.Flights))
	for _, f := range r.Flights {
		flights = append(flights, FlightHit{
			ID: f.ID, Airline: "Vueling", FlightNumber: f.FlightNumber,
			Origin: f.Origin, Destination: f.Destination,
			Depart: f.Departure, Arrive: f.Arrival, Duration: f.Duration,
			Price: fmt.Sprintf("%.2f", f.Price), Currency: f.Currency,
			BookingURL: fmt.Sprintf("%s/es", base),
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
		Brand: brand, Source: "bit",
	}
}
