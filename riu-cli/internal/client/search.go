package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	"github.com/fbelchi/travelkit/parse"
	tkbase "github.com/fbelchi/travelkit/base"
)

const bffHotelsPath = "/api/riucom-bff/v1/hotels"

// Search finds hotels via RIU BFF API or destination ng-state HTML fallback.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)

	html, err := c.fetchDestinationHTML(query)
	if err == nil {
		rows := parse.ParseRIUNgState(html, c.BaseURL, "")
		if len(rows) > 0 {
			return c.rowsToResult(rows, query, page, pageSize, "ng-state")
		}
	}

	if res, err := c.searchBFF(query, page, pageSize); err == nil && len(res.Hotels) > 0 {
		if c.Brand != "" {
			res.Brand = c.Brand
		}
		return res, nil
	}
	return nil, fmt.Errorf("search %q: no hotels — try RIU_COOKIE or `riu session chrome`", query)
}

func (c *Client) rowsToResult(rows []parse.RIUHotel, query string, page, pageSize int, source string) (*HotelSearchResult, error) {
	total := len(rows)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	hits := make([]HotelHit, 0, end-start)
	for _, h := range rows[start:end] {
		price := ""
		if h.Price > 0 {
			price = fmt.Sprintf("%.2f", h.Price)
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: brandOrDefault(c.Brand),
			City: h.City, Country: h.Country, Stars: h.Stars,
			Price: price, Currency: h.Currency, Rating: h.Rating, ReviewCount: h.ReviewCount,
			HotelURL: tkbase.Absolutize(c.BaseURL, "/es/hotel/"+h.Slug),
			ImageURL: h.ImageURL,
		})
	}
	return &HotelSearchResult{
		Query: query, Total: total, Page: page, PageSize: pageSize,
		HasNext: total > page*pageSize, Hotels: hits, Brand: c.Brand, Source: source,
	}, nil
}

func (c *Client) searchBFF(query string, page, pageSize int) (*HotelSearchResult, error) {
	url := fmt.Sprintf("%s%s?offset=%d&limit=%d&has_tripadvisor=false&has_continent=true&fields=id,name,country,destination,slug,price,stars,trip_advisor",
		c.BaseURL, bffHotelsPath, (page-1)*pageSize, pageSize)
	body, status, err := c.GetRaw(url)
	if err != nil {
		return nil, err
	}
	if akamai.IsDenied(status, string(body)) {
		return nil, fmt.Errorf("bff blocked")
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("bff HTTP %d", status)
	}
	var resp struct {
		Body []struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Slug  string `json:"slug"`
			Stars int    `json:"stars"`
			Price *struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"price"`
			Destination *struct {
				Name string `json:"name"`
			} `json:"destination"`
			Country *struct {
				Name string `json:"name"`
			} `json:"country"`
			TripAdvisor *struct {
				Rating  string `json:"rating"`
				Reviews string `json:"reviews"`
			} `json:"trip_advisor"`
		} `json:"body"`
	}
	if err := jsonUnmarshal(body, &resp); err != nil {
		return nil, err
	}
	q := strings.ToLower(query)
	hits := make([]HotelHit, 0, len(resp.Body))
	for _, h := range resp.Body {
		city := ""
		if h.Destination != nil {
			city = h.Destination.Name
		}
		country := ""
		if h.Country != nil {
			country = h.Country.Name
		}
		if q != "" && !strings.Contains(strings.ToLower(h.Name), q) &&
			!strings.Contains(strings.ToLower(city), q) {
			continue
		}
		price, curr := "", "EUR"
		if h.Price != nil {
			price = fmt.Sprintf("%.2f", h.Price.Amount)
			curr = h.Price.Currency
		}
		hits = append(hits, HotelHit{
			ID:       h.ID,
			Name:     h.Name,
			Brand:    brandOrDefault(c.Brand),
			City:     city,
			Country:  country,
			Stars:    float64(h.Stars),
			Price:    price,
			Currency: curr,
			HotelURL: tkbase.Absolutize(c.BaseURL, "/es/hotel/"+h.Slug),
		})
	}
	return &HotelSearchResult{
		Query: query, Total: len(hits), Page: page, PageSize: pageSize,
		Hotels: hits, Brand: c.Brand, Source: "bff",
	}, nil
}

func (c *Client) fetchDestinationHTML(query string) (string, error) {
	slug := strings.ToLower(strings.ReplaceAll(query, " ", "-"))
	paths := []string{
		"/es/hotels/europa/espana/" + slug,
		"/es/hotels/europa/" + slug,
		"/es/hotels/" + slug,
	}
	var lastErr error
	for _, p := range paths {
		html, err := c.FetchHTML(c.BaseURL + p)
		if err != nil {
			lastErr = err
			continue
		}
		if akamai.IsDenied(403, html) {
			lastErr = fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("riu"))
			continue
		}
		if strings.Contains(html, "ng-state") {
			return html, nil
		}
	}
	if lastErr != nil {
		return "", lastErr
	}
	return "", fmt.Errorf("destination page not found for %q", query)
}

func brandOrDefault(brand string) string {
	if brand != "" {
		return brand
	}
	return "RIU Hotels & Resorts"
}
