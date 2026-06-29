package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
	tkhotel "github.com/fbelchi/travelkit/hotel"
	"github.com/fbelchi/travelkit/parse"
)

// Search queries Iberostar hotel search GraphQL BFF, falling back to directory HTML.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	if res, err := c.searchGraphQL(query, page, pageSize); err == nil && len(res.Hotels) > 0 {
		return res, nil
	} else if err != nil && !shouldFallbackIberostar(err) {
		return nil, err
	}
	return c.searchDirectory(query, page, pageSize)
}

func shouldFallbackIberostar(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "akamai blocked") ||
		strings.Contains(msg, "graphql") ||
		strings.Contains(msg, "http 404") ||
		strings.Contains(msg, "http 403")
}

func (c *Client) searchGraphQL(query string, page, pageSize int) (*HotelSearchResult, error) {
	payload := map[string]any{
		"query": `query SearchHotels($q:String!,$page:Int!,$size:Int!){searchHotels(query:$q,page:$page,size:$size){total hasNext hotels{id name brand city country stars minPrice currency url image}}}`,
		"variables": map[string]any{
			"q": query, "page": page, "size": pageSize,
		},
	}
	type hotelRow struct {
		ID       string  `json:"id"`
		Name     string  `json:"name"`
		Brand    string  `json:"brand"`
		City     string  `json:"city"`
		Country  string  `json:"country"`
		Stars    float64 `json:"stars"`
		MinPrice float64 `json:"minPrice"`
		Currency string  `json:"currency"`
		URL      string  `json:"url"`
		Image    string  `json:"image"`
	}
	type searchBlock struct {
		Total   int        `json:"total"`
		HasNext bool       `json:"hasNext"`
		Hotels  []hotelRow `json:"hotels"`
	}
	var resp struct {
		Data struct {
			SearchHotels searchBlock `json:"searchHotels"`
			HotelSearch  searchBlock `json:"hotelSearch"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}
	if err := c.postGraphQL(payload, &resp); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok {
			if akamai.IsDenied(he.Status, he.Body) || akamai.IsAppNotFoundWithoutSession(he.Status, he.Body) {
				return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("iberostar"))
			}
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	if len(resp.Errors) > 0 {
		return nil, fmt.Errorf("graphql: %s", resp.Errors[0].Message)
	}
	sh := resp.Data.SearchHotels
	if len(sh.Hotels) == 0 {
		sh = resp.Data.HotelSearch
	}
	hits := make([]HotelHit, 0, len(sh.Hotels))
	for _, h := range sh.Hotels {
		price := ""
		if h.MinPrice > 0 {
			price = fmt.Sprintf("%.2f", h.MinPrice)
		}
		b := h.Brand
		if c.Brand != "" {
			b = c.Brand
		}
		hits = append(hits, HotelHit{
			ID: h.ID, Name: h.Name, Brand: b, City: h.City, Country: h.Country,
			Stars: h.Stars, Price: price, Currency: h.Currency,
			HotelURL: tkbase.Absolutize(c.BaseURL, h.URL), ImageURL: tkbase.Absolutize(c.BaseURL, h.Image),
		})
	}
	return &HotelSearchResult{
		Query: query, Total: sh.Total, Page: page, PageSize: pageSize,
		HasNext: sh.HasNext, Hotels: hits, Brand: c.Brand, Source: "graphql",
	}, nil
}

func (c *Client) searchDirectory(query string, page, pageSize int) (*HotelSearchResult, error) {
	parseFn := func(html, base string) []parse.HotelLD {
		return parse.HotelsFromIberostarDirectory(html, base, query)
	}
	rows, err := tkhotel.SpanishHTMLSearch(c.FetchHTML, c.BaseURL, tkhotel.IberostarDirectoryPaths(query), parseFn, query)
	if err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("iberostar"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	filtered := tkhotel.FilterSpanishQuery(rows, query)
	if len(filtered) == 0 {
		return nil, fmt.Errorf("search %q: no hotels in directory — %s", query, akamai.NeedsSessionHint("iberostar"))
	}
	brand := c.Brand
	if brand == "" {
		brand = "Iberostar"
	}
	return tkhotel.LDToResult(filtered, query, page, pageSize, brand, c.BaseURL, "directory"), nil
}
