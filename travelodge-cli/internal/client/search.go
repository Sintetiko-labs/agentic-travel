package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
)

const (
	hotelSearchPath = "/api/v2/hotel"
	imageCDNBase    = "https://media.travelodge.co.uk/image/upload/"
	maxSearchPages  = 50
)

type hotelSearchResponse struct {
	Results []hotelRow `json:"results"`
	Links   []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

type hotelRow struct {
	ID        int    `json:"id"`
	HotelURL  string `json:"hotelUrl"`
	MainImage string `json:"mainImage"`
	Code      string `json:"code"`
	Title     string `json:"title"`
	Rating    *struct {
		Reviews int     `json:"reviews"`
		Score   float64 `json:"score"`
	} `json:"rating"`
	Currency        string  `json:"currency"`
	MinPrice        float64 `json:"minPrice"`
	HasAvailability bool    `json:"hasAvailability"`
}

// Search finds hotels via Travelodge /api/v2/hotel (public JSON API).
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, fmt.Errorf("destination query required")
	}

	checkIn, checkOut := defaultStayDates()
	rows, err := c.fetchAllHotels(query, checkIn, checkOut)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("search %q: no hotels found", query)
	}

	total := len(rows)
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}

	hits := make([]HotelHit, 0, end-start)
	for _, h := range rows[start:end] {
		hits = append(hits, rowToHit(c, h, query))
	}

	return &HotelSearchResult{
		Query:    query,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		HasNext:  total > page*pageSize,
		Hotels:   hits,
		Brand:    brandOrDefault(c.Brand),
		Source:   "api/v2/hotel",
	}, nil
}

func (c *Client) fetchAllHotels(query, checkIn, checkOut string) ([]hotelRow, error) {
	var all []hotelRow
	start := 0
	for page := 0; page < maxSearchPages; page++ {
		resp, err := c.fetchHotelPage(query, checkIn, checkOut, start)
		if err != nil {
			if len(all) > 0 {
				return all, nil
			}
			return nil, err
		}
		if len(resp.Results) == 0 {
			break
		}
		all = append(all, resp.Results...)
		if !hasNextLink(resp.Links) {
			break
		}
		start += len(resp.Results)
	}
	return all, nil
}

func (c *Client) fetchHotelPage(query, checkIn, checkOut string, start int) (*hotelSearchResponse, error) {
	q := url.Values{}
	q.Set("pagination", "false")
	q.Set("checkIn", checkIn)
	q.Set("checkOut", checkOut)
	q.Set("q", query)
	q.Set("rooms[0][adults]", "1")
	q.Set("rooms[0][children]", "0")
	q.Set("action", "search")
	if start > 0 {
		q.Set("start", strconv.Itoa(start))
	}

	rawURL := c.BaseURL + hotelSearchPath + "?" + q.Encode()
	body, status, err := c.GetRaw(rawURL)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		return nil, &APIError{Status: status, Body: tkbase.Truncate(string(body), 300)}
	}

	var resp hotelSearchResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("decode hotel search: %w", err)
	}
	return &resp, nil
}

func rowToHit(c *Client, h hotelRow, query string) HotelHit {
	price := ""
	if h.MinPrice > 0 {
		price = fmt.Sprintf("%.2f", h.MinPrice)
	}
	id := strconv.Itoa(h.ID)
	if h.Code != "" {
		id = h.Code
	}
	hotelURL := cleanHotelURL(c.BaseURL, h.HotelURL)
	rating := 0.0
	reviews := 0
	if h.Rating != nil {
		rating = h.Rating.Score
		reviews = h.Rating.Reviews
	}
	return HotelHit{
		ID:          id,
		Name:        h.Title,
		Brand:       brandOrDefault(c.Brand),
		City:        query,
		Country:     "United Kingdom",
		Price:       price,
		Currency:    h.Currency,
		Rating:      rating,
		ReviewCount: reviews,
		HotelURL:    hotelURL,
		ImageURL:    hotelImageURL(h.MainImage),
	}
}

func cleanHotelURL(base, raw string) string {
	if raw == "" {
		return ""
	}
	u, err := url.Parse(raw)
	if err != nil {
		return tkbase.Absolutize(base, raw)
	}
	u.RawQuery = ""
	u.Fragment = ""
	return tkbase.Absolutize(base, u.Path)
}

func hotelImageURL(path string) string {
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	return imageCDNBase + strings.TrimPrefix(path, "/")
}

func hasNextLink(links []struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}) bool {
	for _, l := range links {
		if l.Rel == "next" && l.Href != "" {
			return true
		}
	}
	return false
}

func defaultStayDates() (checkIn, checkOut string) {
	return "2026-07-05", "2026-07-06"
}

func brandOrDefault(brand string) string {
	if brand != "" {
		return brand
	}
	return "Travelodge"
}
