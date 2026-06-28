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
		return nil, fmt.Errorf("bootstrap: %w", err)
	}
	body, status, err := c.postPricing(buildJourneyPricePayload(origin, dest, depart, ret))
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w — try plusultra session chrome --wait", origin, dest, err)
	}
	if akamai.IsWAFBlocked(status, string(body)) {
		return nil, fmt.Errorf("waf blocked — %s", akamai.NeedsSessionHint("plusultra"))
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("search %s→%s: HTTP %d: %s", origin, dest, status, tkbase.Truncate(string(body), 200))
	}
	flights, err := parsePlusUltra(body, origin, dest, depart)
	if err != nil {
		return nil, err
	}
	return paginate(origin, dest, depart, ret, page, pageSize, flights, c.Brand, "pricing/journeyPrice"), nil
}

func buildJourneyPricePayload(origin, dest, depart, ret string) map[string]any {
	d, _ := time.Parse("2006-01-02", depart)
	reqs := []map[string]any{{
		"id": "1", "origin": origin, "destination": dest, "currency": "EUR",
		"pax":     map[string]int{"ADT": 1, "CHD": 0, "INF": 0, "TNG": 0},
		"details": map[string]any{"default": []map[string]string{{"begin": d.AddDate(0, 0, -9).Format("2006-01-02"), "end": d.AddDate(0, 0, 10).Format("2006-01-02")}}},
	}}
	if ret != "" {
		rd, _ := time.Parse("2006-01-02", ret)
		reqs = append(reqs, map[string]any{
			"id": "2", "origin": dest, "destination": origin, "currency": "EUR",
			"pax": map[string]int{"ADT": 1}, "details": map[string]any{"default": []map[string]string{{"begin": rd.AddDate(0, 0, -3).Format("2006-01-02"), "end": rd.AddDate(0, 0, 3).Format("2006-01-02")}}},
		})
	}
	return map[string]any{"customerId": CustomerGUID, "journeyPriceRequests": reqs}
}

func (c *Client) postPricing(payload map[string]any) ([]byte, int, error) {
	b, _ := json.Marshal(payload)
	c.Throttle()
	req, _ := http.NewRequest(http.MethodPost, PricingAPI+"/journeyPrice", strings.NewReader(string(b)))
	c.SetAPIHeaders(req)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("x-api-key", PricingKey)
	req.Header.Set("origin", BaseURL)
	req.Header.Set("referer", BaseURL+"/es-es/")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	return raw, resp.StatusCode, nil
}

func parsePlusUltra(raw []byte, origin, dest, depart string) ([]FlightHit, error) {
	var resp struct {
		JourneyPriceResponses []struct {
			Currency  string `json:"currency"`
			Schedules []struct {
				Date     string `json:"date"`
				Journeys []struct {
					Origin, Destination, Std, Sta string
					TotalAmount                   float64 `json:"totalAmount"`
					Segments                      []struct{ FlightNumber string `json:"flightNumber"` } `json:"segments"`
				} `json:"journeys"`
			} `json:"schedules"`
		} `json:"journeyPriceResponses"`
	}
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	out := []FlightHit{}
	for _, jpr := range resp.JourneyPriceResponses {
		curr := tkbase.FirstNonEmpty(jpr.Currency, "EUR")
		for _, s := range jpr.Schedules {
			for _, j := range s.Journeys {
				day := strings.Split(tkbase.FirstNonEmpty(j.Std, s.Date), "T")[0]
				if day != depart && s.Date != depart {
					continue
				}
				if strings.ToUpper(j.Origin) != origin || strings.ToUpper(j.Destination) != dest {
					continue
				}
				fn := "PU"
				if len(j.Segments) > 0 && j.Segments[0].FlightNumber != "" {
					fn = j.Segments[0].FlightNumber
				}
				out = append(out, FlightHit{ID: fn + "-" + depart, Airline: "Plus Ultra", FlightNumber: fn, Origin: origin, Destination: dest, Depart: isoTime(j.Std), Arrive: isoTime(j.Sta), Price: fmt.Sprintf("%.2f", j.TotalAmount), Currency: curr, BookingURL: BaseURL + "/es-es/"})
			}
		}
	}
	return out, nil
}

func isoTime(s string) string {
	if i := strings.Index(s, "T"); i >= 0 && len(s) >= i+6 {
		return s[i+1 : i+6]
	}
	return s
}

func paginate(origin, dest, depart, ret string, page, pageSize int, flights []FlightHit, brand, source string) *FlightSearchResult {
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
	return &FlightSearchResult{Query: origin + "-" + dest + " " + depart, Origin: origin, Dest: dest, Depart: depart, Return: ret, Total: total, Page: page, PageSize: pageSize, HasNext: total > page*pageSize, Flights: pageFlights, Brand: brand, Source: source}
}
