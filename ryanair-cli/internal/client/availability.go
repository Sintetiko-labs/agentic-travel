package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	tkbase "github.com/fbelchi/travelkit/base"
)

const bookingLocale = "es-es"

var clientVersionRE = regexp.MustCompile(`client-version["'\s,:=]+([0-9]+\.[0-9]+\.[0-9]+)`)

func (c *Client) ensureBookingSession() {
	u := c.BaseURL + "/"
	if c.Jar == nil {
		return
	}
	cookies := c.Jar.Cookies(mustURL(u))
	for _, ck := range cookies {
		if ck.Name == "fr-correlation-id" {
			return
		}
	}
	c.Jar.SetCookies(mustURL(u), []*http.Cookie{{
		Name:  "fr-correlation-id",
		Value: fmt.Sprintf("travel-cli-%d", time.Now().Unix()),
		Path:  "/",
	}})
}

func (c *Client) clientVersion() string {
	if v := strings.TrimSpace(os.Getenv("RYANAIR_CLIENT_VERSION")); v != "" {
		return v
	}
	if c.clientVer != "" {
		return c.clientVer
	}
	html, err := c.FetchHTML(c.BaseURL + "/es/es")
	if err == nil {
		if m := clientVersionRE.FindStringSubmatch(html); len(m) > 1 {
			c.clientVer = m[1]
			return c.clientVer
		}
	}
	c.clientVer = "3.197.0"
	return c.clientVer
}

func (c *Client) refreshClientVersion() {
	c.clientVer = ""
	_ = c.clientVersion()
}

func (c *Client) searchAvailability(origin, dest, depart, ret string) (*FlightSearchResult, error) {
	c.ensureBookingSession()
	ver := c.clientVersion()
	path := fmt.Sprintf("/api/booking/v4/%s/availability?ADT=1&CHD=0&DateIn=%s&DateOut=%s&Destination=%s&Disc=0&INF=0&Origin=%s&TEEN=0&FlexDaysIn=2&FlexDaysBeforeIn=2&FlexDaysOut=2&FlexDaysBeforeOut=2&ToUs=AGREED&IncludeConnectingFlights=false&RoundTrip=%t",
		bookingLocale, ret, depart, dest, origin, ret != "")
	body, status, err := c.getBooking(path, ver)
	if err != nil {
		return nil, err
	}
	if status == http.StatusConflict && strings.Contains(string(body), "Availability declined") {
		c.refreshClientVersion()
		body, status, err = c.getBooking(path, c.clientVersion())
		if err != nil {
			return nil, err
		}
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("availability API HTTP %d: %s", status, tkbase.Truncate(string(body), 200))
	}
	return parseAvailabilityJSON(body, origin, dest, depart, ret)
}

func (c *Client) getBooking(path, ver string) ([]byte, int, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return nil, 0, err
	}
	c.SetAPIHeaders(req)
	req.Header.Set("client-version", ver)
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	return body, resp.StatusCode, nil
}

func (c *Client) searchFarfnd(origin, dest, depart string, page, pageSize int) (*FlightSearchResult, error) {
	month := depart
	if len(month) >= 7 {
		month = month[:7] + "-01"
	}
	path := fmt.Sprintf("/api/farfnd/v4/oneWayFares/%s/%s/cheapestPerDay?outboundMonthOfDate=%s", origin, dest, month)
	body, status, err := c.GetRaw(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("farfnd HTTP %d", status)
	}
	return parseFarfndResponse(body, c.BaseURL, origin, dest, depart, page, pageSize)
}

func parseFarfndResponse(body []byte, baseURL, origin, dest, depart string, page, pageSize int) (*FlightSearchResult, error) {
	var raw farfndResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}
	return buildFarfndResult(&raw, baseURL, origin, dest, depart, page, pageSize), nil
}

func buildFarfndResult(raw *farfndResponse, baseURL, origin, dest, depart string, page, pageSize int) *FlightSearchResult {
	all := make([]FlightHit, 0, len(raw.Outbound.Fares))
	day := depart
	for _, f := range raw.Outbound.Fares {
		if f.Unavailable || f.Price == nil {
			continue
		}
		if day != "" && f.Day != day {
			continue
		}
		all = append(all, FlightHit{
			ID:          fmt.Sprintf("%s-%s-%s", origin, dest, f.Day),
			Airline:     "Ryanair",
			Origin:      origin,
			Destination: dest,
			Depart:      firstNonEmpty(f.DepartureDate, f.Day+"T00:00:00"),
			Arrive:      f.ArrivalDate,
			Price:       fmt.Sprintf("%.2f", f.Price.Value),
			Currency:    f.Price.CurrencyCode,
			BookingURL:  fmt.Sprintf("%s/es/es/cheap-flights/%s-to-%s?outboundDate=%s", baseURL, origin, dest, f.Day),
		})
	}
	total := len(all)
	start := (page - 1) * pageSize
	var pageHits []FlightHit
	if start < total {
		end := start + pageSize
		if end > total {
			end = total
		}
		pageHits = all[start:end]
	}
	return &FlightSearchResult{
		Query:    fmt.Sprintf("%s-%s %s", origin, dest, depart),
		Origin:   origin,
		Dest:     dest,
		Depart:   depart,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		HasNext:  total > page*pageSize,
		Flights:  pageHits,
		Source:   "farfnd",
	}
}

type farfndResponse struct {
	Outbound struct {
		Fares []struct {
			Day           string `json:"day"`
			ArrivalDate   string `json:"arrivalDate"`
			DepartureDate string `json:"departureDate"`
			Unavailable   bool   `json:"unavailable"`
			Price         *struct {
				Value        float64 `json:"value"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"price"`
		} `json:"fares"`
	} `json:"outbound"`
}

func parseAvailabilityJSON(raw []byte, origin, dest, depart, ret string) (*FlightSearchResult, error) {
	var resp struct {
		Trips []struct {
			Dates []struct {
				Flights []struct {
					FlightNumber string   `json:"flightNumber"`
					Time         []string `json:"time"`
					Fares        []struct {
						Amount float64 `json:"amount"`
					} `json:"fares"`
				} `json:"flights"`
			} `json:"dates"`
		} `json:"trips"`
	}
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	flights := make([]FlightHit, 0)
	for _, trip := range resp.Trips {
		for _, d := range trip.Dates {
			for _, fl := range d.Flights {
				price := ""
				curr := "EUR"
				if len(fl.Fares) > 0 {
					price = fmt.Sprintf("%.2f", fl.Fares[0].Amount)
				}
				depAt := depart
				if len(fl.Time) > 0 {
					depAt = fl.Time[0]
				}
				flights = append(flights, FlightHit{
					ID:           strings.ReplaceAll(fl.FlightNumber, " ", "-"),
					Airline:      "Ryanair",
					FlightNumber: fl.FlightNumber,
					Origin:       origin,
					Destination:  dest,
					Depart:       depAt,
					Price:        price,
					Currency:     curr,
					BookingURL:   fmt.Sprintf("https://www.ryanair.com/es/es/cheap-flights/%s-to-%s", origin, dest),
				})
			}
		}
	}
	return &FlightSearchResult{
		Query:    fmt.Sprintf("%s-%s %s", origin, dest, depart),
		Origin:   origin,
		Dest:     dest,
		Depart:   depart,
		Return:   ret,
		Total:    len(flights),
		Page:     1,
		PageSize: len(flights),
		Flights:  flights,
		Source:   "booking",
	}, nil
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

func mustURL(s string) *url.URL {
	u, _ := url.Parse(s)
	return u
}
