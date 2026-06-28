package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search queries Iberia Express via Iberia booking API (carrier I2).
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

	q := fmt.Sprintf("/api/availability/v1/flights?market=ES&language=es&origin=%s&destination=%s&departureDate=%s&adults=1&operatingCarrier=I2",
		origin, dest, depart)
	if ret != "" {
		q += "&returnDate=" + ret
	}
	var resp iberiaExpressResponse
	if err := c.GetJSON(q, &resp); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("incapsula blocked — %s", akamai.NeedsSessionHint("iberiaexpress"))
		}
		return nil, fmt.Errorf("search %s→%s: %w — try IBERIAEXPRESS_COOKIE", origin, dest, err)
	}
	return resp.toResult(origin, dest, depart, ret, page, pageSize, c.Brand), nil
}

type iberiaExpressResponse struct {
	Offers []struct {
		ID           string  `json:"id"`
		FlightNumber string  `json:"flightNumber"`
		Origin       string  `json:"origin"`
		Destination  string  `json:"destination"`
		Departure    string  `json:"departure"`
		Arrival      string  `json:"arrival"`
		Price        float64 `json:"price"`
		Currency     string  `json:"currency"`
	} `json:"offers"`
}

func (r *iberiaExpressResponse) toResult(origin, dest, depart, ret string, page, pageSize int, brand string) *FlightSearchResult {
	flights := make([]FlightHit, 0, len(r.Offers))
	for _, f := range r.Offers {
		flights = append(flights, FlightHit{
			ID: f.ID, Airline: "Iberia Express", FlightNumber: f.FlightNumber,
			Origin: f.Origin, Destination: f.Destination,
			Depart: f.Departure, Arrive: f.Arrival,
			Price: fmt.Sprintf("%.2f", f.Price), Currency: f.Currency,
			BookingURL: cBaseURL + "/es/",
		})
	}
	total := len(flights)
	return &FlightSearchResult{
		Query: fmt.Sprintf("%s-%s %s", origin, dest, depart),
		Origin: origin, Dest: dest, Depart: depart, Return: ret,
		Total: total, Page: page, PageSize: pageSize,
		Flights: flights, Brand: brand, Source: "api",
	}
}

const cBaseURL = "https://www.iberiaexpress.com"
