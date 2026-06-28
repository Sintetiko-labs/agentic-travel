package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
)

// Search queries easyJet ejavailability API.
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

	_, _ = c.FetchHTML(c.BaseURL + "/es/")
	q := url.Values{}
	q.Set("AdditionalSeats", "0")
	q.Set("AdultSeats", "1")
	q.Set("ArrivalIata", dest)
	q.Set("ChildSeats", "0")
	q.Set("DepartureIata", origin)
	q.Set("IncludeAdminFees", "true")
	q.Set("IncludeFlexiFares", "false")
	q.Set("IncludeLowestFareSeats", "true")
	q.Set("IncludePrices", "true")
	q.Set("Infants", "0")
	q.Set("IsTransfer", "false")
	q.Set("LanguageCode", "ES")
	q.Set("MinDepartureDate", depart)
	q.Set("MaxDepartureDate", depart)
	if ret != "" {
		q.Set("MinReturnDate", ret)
		q.Set("MaxReturnDate", ret)
	}
	apiURL := c.BaseURL + "/ejavailability/api/v5/availability/query?" + q.Encode()
	body, status, err := c.getAvailability(apiURL)
	if err != nil {
		return nil, err
	}
	if akamai.IsWAFBlocked(status, string(body)) {
		return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("easyjet"))
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("availability HTTP %d", status)
	}
	if !akamai.LooksLikeJSON(string(body)) {
		return nil, fmt.Errorf("availability: non-JSON response — %s", akamai.NeedsSessionHint("easyjet"))
	}
	var resp easyjetResponse
	if err := jsonUnmarshal(body, &resp); err != nil {
		return nil, err
	}
	return resp.toResult(origin, dest, depart, ret, page, pageSize, c.Brand), nil
}

func (c *Client) getAvailability(fullURL string) ([]byte, int, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
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

type easyjetResponse struct {
	AvailableFlights []struct {
		CarrierCode    string `json:"CarrierCode"`
		FlightNumber   int    `json:"FlightNumber"`
		DepartureIata  string `json:"DepartureIata"`
		ArrivalIata    string `json:"ArrivalIata"`
		LocalDeparture string `json:"LocalDepartureTime"`
		LocalArrival   string `json:"LocalArrivalTime"`
		FlightFares    []struct {
			Prices struct {
				Adult struct {
					Price float64 `json:"Price"`
				} `json:"Adult"`
			} `json:"Prices"`
		} `json:"FlightFares"`
	} `json:"AvailableFlights"`
}

func (r *easyjetResponse) toResult(origin, dest, depart, ret string, page, pageSize int, brand string) *FlightSearchResult {
	flights := make([]FlightHit, 0, len(r.AvailableFlights))
	for _, f := range r.AvailableFlights {
		price := ""
		curr := "EUR"
		if len(f.FlightFares) > 0 {
			price = fmt.Sprintf("%.2f", f.FlightFares[0].Prices.Adult.Price)
		}
		fn := fmt.Sprintf("%s %d", f.CarrierCode, f.FlightNumber)
		flights = append(flights, FlightHit{
			ID: fmt.Sprintf("%s-%s-%s", f.DepartureIata, f.ArrivalIata, depart),
			Airline: "easyJet", FlightNumber: fn,
			Origin: f.DepartureIata, Destination: f.ArrivalIata,
			Depart: f.LocalDeparture, Arrive: f.LocalArrival,
			Price: price, Currency: curr,
			BookingURL: "https://www.easyjet.com/es/",
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
		Brand: brand, Source: "ejavailability",
	}
}
