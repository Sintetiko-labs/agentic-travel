package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

const (
	skysalesBaseURL  = "https://tickets.vueling.com"
	flightPriceAPI   = "https://apiwww.vueling.com/api/FlightPrice/GetAllFlights"
)

// Search queries Vueling public FlightPrice calendar API (apiwww.vueling.com).
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

	departTime, err := time.Parse("2006-01-02", depart)
	if err != nil {
		return nil, fmt.Errorf("search %s→%s: invalid depart date %q (use YYYY-MM-DD)", origin, dest, depart)
	}

	q := url.Values{}
	q.Set("originCode", origin)
	q.Set("destinationCode", dest)
	q.Set("year", fmt.Sprintf("%d", departTime.Year()))
	q.Set("month", fmt.Sprintf("%d", int(departTime.Month())))
	q.Set("currencyCode", "EUR")
	q.Set("monthsRange", "1")

	var rows []vuelingFlightPriceRow
	err = c.getJSON(flightPriceAPI+"?"+q.Encode(), &rows)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("vueling"))
		}
		return nil, fmt.Errorf("search %s→%s: %w", origin, dest, err)
	}

	flights := filterVuelingRows(rows, origin, dest, depart, ret)
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
		Source:   "flightprice",
	}, nil
}

func (c *Client) getJSON(fullURL string, out any) error {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return err
	}
	c.SetAPIHeaders(req)
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("origin", BaseURL)
	req.Header.Set("referer", BaseURL+"/")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &tkbase.HTTPError{Status: resp.StatusCode, Body: tkbase.Truncate(string(body), 300)}
	}
	if err := jsonUnmarshal(body, out); err != nil {
		return err
	}
	return nil
}

type vuelingFlightPriceRow struct {
	ArrivalDate      string  `json:"ArrivalDate"`
	ArrivalStation   string  `json:"ArrivalStation"`
	DepartureDate    string  `json:"DepartureDate"`
	DepartureStation string  `json:"DepartureStation"`
	FlightID         string  `json:"FlightID"`
	Price            float64 `json:"Price"`
	IsInvalidPrice   bool    `json:"IsInvalidPrice"`
}

func filterVuelingRows(rows []vuelingFlightPriceRow, origin, dest, depart, ret string) []FlightHit {
	departDay := depart
	var out []FlightHit
	for _, r := range rows {
		if r.IsInvalidPrice || r.Price <= 0 {
			continue
		}
		if !strings.EqualFold(r.DepartureStation, origin) || !strings.EqualFold(r.ArrivalStation, dest) {
			continue
		}
		dep := strings.Split(r.DepartureDate, "T")[0]
		if dep != departDay {
			continue
		}
		arr := r.ArrivalDate
		if i := strings.Index(arr, "T"); i > 0 {
			arr = arr[i+1:]
			if len(arr) >= 5 {
				arr = arr[:5]
			}
		}
		depTime := r.DepartureDate
		if i := strings.Index(depTime, "T"); i > 0 {
			depTime = depTime[i+1:]
			if len(depTime) >= 5 {
				depTime = depTime[:5]
			}
		}
		fn := "VY" + strings.TrimLeft(r.FlightID, "0")
		if strings.HasPrefix(strings.ToUpper(r.FlightID), "VY") {
			fn = strings.ToUpper(r.FlightID)
		}
		out = append(out, FlightHit{
			ID:           fmt.Sprintf("%s-%s-%s-%s", origin, dest, dep, r.FlightID),
			Airline:      "Vueling",
			FlightNumber: fn,
			Origin:       origin,
			Destination:  dest,
			Depart:       depTime,
			Arrive:       arr,
			Price:        fmt.Sprintf("%.2f", r.Price),
			Currency:     "EUR",
			BookingURL:   vuelingBookingURL(origin, dest, depart, ret, fn),
		})
	}
	return out
}

func vuelingBookingURL(origin, dest, depart, ret, flightNum string) string {
	q := url.Values{}
	q.Set("o", origin)
	q.Set("d", dest)
	q.Set("dd", depart)
	if ret != "" {
		q.Set("rd", ret)
		q.Set("dt", "2")
	} else {
		q.Set("dt", "1")
	}
	q.Set("adt", "1")
	q.Set("c", "es-ES")
	q.Set("cur", "EUR")
	if flightNum != "" {
		q.Set("ofn", flightNum)
	}
	return skysalesBaseURL + "/booking?" + q.Encode()
}
