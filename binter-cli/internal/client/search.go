package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

const binterGraphQL = "https://services.bintercanarias.com/main/graphql"

const bookingSearchQuery = `query booking_search($search: BookingSearchInput!, $language: LanguageInput) {
  booking_search(search: $search, language: $language) {
    availability {
      flights {
        flightSegments {
          line
          flightNumber
          origin { iata }
          destination { iata }
          departureTime
          arrivalTime
          stops
        }
        fares {
          bestPrice
          amount { totalAmount currency { code } }
        }
      }
    }
  }
}`

// Search queries Binter GraphQL booking_search API.
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

	if _, err := c.FetchHTML(BaseURL + "/es"); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("cloudflare blocked — %s", akamai.NeedsSessionHint("binter"))
		}
	}

	hash1, err := c.binterHash1()
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: hash: %w", origin, dest, err)
	}

	variables := map[string]any{
		"search":   binterSearchInput(origin, dest, depart, ret),
		"language": "es",
	}

	body, status, err := c.binterGraphQL("booking_search", bookingSearchQuery, variables, hash1)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}
	if status < 200 || status >= 300 {
		if akamai.IsDenied(status, string(body)) {
			return nil, fmt.Errorf("cloudflare blocked — %s", akamai.NeedsSessionHint("binter"))
		}
		return nil, fmt.Errorf("search %s→%s: HTTP %d: %s", origin, dest, status, tkbase.Truncate(string(body), 200))
	}

	flights, gqlErr := parseBinterSearch(body, origin, dest, depart)
	if gqlErr != "" && len(flights) == 0 {
		return nil, fmt.Errorf("search %s→%s: %s", origin, dest, gqlErr)
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
		Source:   "booking_search",
	}, nil
}

func binterSearchInput(origin, dest, depart, ret string) map[string]any {
	trips := []map[string]any{{"origin": origin, "destination": dest, "date": depart}}
	if ret != "" {
		trips = append(trips, map[string]any{"origin": dest, "destination": origin, "date": ret})
	}
	return map[string]any{
		"trips": trips,
		"passengers": []map[string]any{{
			"type":     map[string]any{"code": "ADT"},
			"quantity": 1,
		}},
	}
}

func (c *Client) binterHash1() (string, error) {
	body, status, err := c.binterGraphQL("hash_getHash1", `query hash_getHash1 { hash_getHash1 }`, nil, "")
	if err != nil {
		return "", err
	}
	if status < 200 || status >= 300 {
		return "", fmt.Errorf("HTTP %d: %s", status, tkbase.Truncate(string(body), 200))
	}
	var out struct {
		Data struct {
			Hash string `json:"hash_getHash1"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
			Code    string `json:"code"`
		} `json:"errors"`
	}
	if err := jsonUnmarshal(body, &out); err != nil {
		return "", err
	}
	if out.Data.Hash != "" {
		return out.Data.Hash, nil
	}
	if len(out.Errors) > 0 {
		return "", fmt.Errorf("%s (%s)", out.Errors[0].Message, out.Errors[0].Code)
	}
	return "", fmt.Errorf("empty hash_getHash1")
}

func (c *Client) binterGraphQL(op, query string, variables map[string]any, hash1 string) ([]byte, int, error) {
	if variables == nil {
		variables = map[string]any{}
	}
	payload, err := json.Marshal(map[string]any{
		"operationName": op,
		"query":         query,
		"variables":     variables,
	})
	if err != nil {
		return nil, 0, err
	}
	c.Throttle()
	req, err := http.NewRequest(http.MethodPost, binterGraphQL+"?"+op, bytes.NewReader(payload))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("origin", BaseURL)
	req.Header.Set("referer", BaseURL+"/es")
	req.Header.Set("x-environment", "desktop")
	req.Header.Set("x-lang", "es")
	if hash1 != "" {
		req.Header.Set("x-hash1", hash1)
	}
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	return raw, resp.StatusCode, nil
}

type binterSearchResponse struct {
	Data struct {
		BookingSearch struct {
			Availability []struct {
				Flights []binterFlight `json:"flights"`
			} `json:"availability"`
		} `json:"booking_search"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type binterFlight struct {
	FlightSegments []struct {
		Line          string `json:"line"`
		FlightNumber  int    `json:"flightNumber"`
		Origin        struct{ IATA string `json:"iata"` } `json:"origin"`
		Destination   struct{ IATA string `json:"iata"` } `json:"destination"`
		DepartureTime string `json:"departureTime"`
		ArrivalTime   string `json:"arrivalTime"`
		Stops         int    `json:"stops"`
	} `json:"flightSegments"`
	Fares []binterFare `json:"fares"`
}

type binterFare struct {
	BestPrice bool `json:"bestPrice"`
	Amount    struct {
		TotalAmount float64 `json:"totalAmount"`
		Currency    struct {
			Code string `json:"code"`
		} `json:"currency"`
	} `json:"amount"`
}

func parseBinterSearch(raw []byte, origin, dest, depart string) ([]FlightHit, string) {
	var resp binterSearchResponse
	if err := jsonUnmarshal(raw, &resp); err != nil {
		return []FlightHit{}, err.Error()
	}
	errMsg := ""
	if len(resp.Errors) > 0 {
		errMsg = resp.Errors[0].Message
	}
	return filterBinterFlights(resp, origin, dest, depart), errMsg
}

func filterBinterFlights(resp binterSearchResponse, origin, dest, depart string) []FlightHit {
	out := make([]FlightHit, 0)
	for _, avail := range resp.Data.BookingSearch.Availability {
		for _, f := range avail.Flights {
			if len(f.FlightSegments) == 0 {
				continue
			}
			seg := f.FlightSegments[0]
			if strings.Split(seg.DepartureTime, "T")[0] != depart {
				continue
			}
			if strings.ToUpper(seg.Origin.IATA) != origin || strings.ToUpper(seg.Destination.IATA) != dest {
				continue
			}
			price, curr := binterBestFare(f.Fares)
			fn := binterFlightNumber(seg.Line, seg.FlightNumber)
			out = append(out, FlightHit{
				ID:           fmt.Sprintf("%s-%s-%s-%s", origin, dest, depart, fn),
				Airline:      "Binter",
				FlightNumber: fn,
				Origin:       origin,
				Destination:  dest,
				Depart:       binterTime(seg.DepartureTime),
				Arrive:       binterTime(seg.ArrivalTime),
				Stops:        seg.Stops,
				Price:        price,
				Currency:     curr,
				BookingURL:   binterBookingURL(origin, dest, depart),
			})
		}
	}
	return out
}

func binterBestFare(fares []binterFare) (string, string) {
	var best *binterFare
	for i := range fares {
		f := &fares[i]
		if f.Amount.Currency.Code != "EUR" {
			continue
		}
		if f.BestPrice || best == nil {
			best = f
			if f.BestPrice {
				break
			}
		}
	}
	if best == nil {
		return "", "EUR"
	}
	return fmt.Sprintf("%.2f", best.Amount.TotalAmount), best.Amount.Currency.Code
}

func binterFlightNumber(line string, num int) string {
	line = strings.TrimSpace(strings.ToUpper(line))
	if line != "" {
		return fmt.Sprintf("%s%d", line, num)
	}
	return fmt.Sprintf("NT%d", num)
}

func binterTime(iso string) string {
	if t, err := time.Parse(time.RFC3339, iso); err == nil {
		return t.Format("15:04")
	}
	if i := strings.Index(iso, "T"); i >= 0 && len(iso) >= i+6 {
		return iso[i+1 : i+6]
	}
	return iso
}

func binterBookingURL(origin, dest, depart string) string {
	q := url.Values{}
	q.Set("origin", origin)
	q.Set("destination", dest)
	q.Set("departure", depart)
	return BaseURL + "/es/reservas?" + q.Encode()
}
