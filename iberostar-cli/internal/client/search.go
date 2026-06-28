package client

import (
	"fmt"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
)

// Search queries Iberostar hotel search GraphQL BFF.
func (c *Client) Search(query string, page, pageSize int) (*HotelSearchResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 24
	}
	query = strings.TrimSpace(query)
	payload := map[string]any{
		"query": `query Search($q:String!,$page:Int!,$size:Int!){searchHotels(query:$q,page:$page,size:$size){total hasNext hotels{id name brand city country stars minPrice currency url image}}}`,
		"variables": map[string]any{
			"q": query, "page": page, "size": pageSize,
		},
	}
	var resp struct {
		Data struct {
			SearchHotels struct {
				Total   int  `json:"total"`
				HasNext bool `json:"hasNext"`
				Hotels  []struct {
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
				} `json:"hotels"`
			} `json:"searchHotels"`
		} `json:"data"`
	}
	if err := c.PostJSON(c.BaseURL+"/api/graphql", payload, &resp); err != nil {
		if he, ok := err.(*tkbase.HTTPError); ok && akamai.IsDenied(he.Status, he.Body) {
			return nil, fmt.Errorf("akamai blocked — %s", akamai.NeedsSessionHint("iberostar"))
		}
		return nil, fmt.Errorf("search %q: %w", query, err)
	}
	sh := resp.Data.SearchHotels
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
