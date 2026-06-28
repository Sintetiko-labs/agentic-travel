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

const availabilityPath = "/api/availability/v1/flights"

// Search queries Iberia Express availability API (Incapsula-protected).
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

	if _, err := c.FetchHTML(BaseURL + "/es/"); err != nil {
		return nil, fmt.Errorf("bootstrap session: %w", err)
	}

	q := fmt.Sprintf("%s?market=ES&language=es&origin=%s&destination=%s&departureDate=%s&adults=1&operatingCarrier=I2",
		availabilityPath, origin, dest, depart)
	if ret != "" {
		q += "&returnDate=" + ret
	}

	body, status, err := c.getAvailability(q)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w — try iberiaexpress session chrome --wait", origin, dest, err)
	}
	if akamai.IsWAFBlocked(status, string(body)) {
		return nil, fmt.Errorf("incapsula blocked — %s", akamai.NeedsSessionHint("iberiaexpress"))
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("search %s→%s: HTTP %d: %s", origin, dest, status, tkbase.Truncate(string(body), 200))
	}
	if !akamai.LooksLikeJSON(string(body)) {
		return nil, fmt.Errorf("search %s→%s: non-JSON response — %s", origin, dest, akamai.NeedsSessionHint("iberiaexpress"))
	}
	var resp iberiaExpressResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("search %s→%s: decode: %w", origin, dest, err)
	}
	return resp.toResult(origin, dest, depart, ret, page, pageSize, c.Brand), nil
}

func (c *Client) getAvailability(path string) ([]byte, int, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return nil, 0, err
	}
	c.SetAPIHeaders(req)
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("origin", c.BaseURL)
	req.Header.Set("referer", c.BaseURL+"/es/")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	return body, resp.StatusCode, nil
}

type iberiaExpressOffer struct {
	ID           string  `json:"id"`
	FlightNumber string  `json:"flightNumber"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	Departure    string  `json:"departure"`
	Arrival      string  `json:"arrival"`
	Price        float64 `json:"price"`
	Currency     string  `json:"currency"`
}

type iberiaExpressResponse struct {
	Offers  []iberiaExpressOffer `json:"offers"`
	Flights []iberiaExpressOffer `json:"flights"`
}

func (r *iberiaExpressResponse) toResult(origin, dest, depart, ret string, page, pageSize int, brand string) *FlightSearchResult {
	items := r.Offers
	if len(items) == 0 {
		items = r.Flights
	}
	flights := make([]FlightHit, 0, len(items))
	for _, f := range items {
		curr := f.Currency
		if curr == "" {
			curr = "EUR"
		}
		flights = append(flights, FlightHit{
			ID: f.ID, Airline: "Iberia Express", FlightNumber: f.FlightNumber,
			Origin: f.Origin, Destination: f.Destination,
			Depart: f.Departure, Arrive: f.Arrival,
			Price: fmt.Sprintf("%.2f", f.Price), Currency: curr,
			BookingURL: BaseURL + "/es/",
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
		Brand: brand, Source: "availability",
	}
}
