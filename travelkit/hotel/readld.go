package hotel

import (
	"fmt"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/parse"
	tktypes "github.com/fbelchi/travelkit/types"
)

// LDReader loads hotel detail from listing search and/or JSON-LD on the hotel page.
type LDReader struct {
	BaseURL   string
	Brand     string
	FetchHTML func(url string) (string, error)
	URLForID  func(id string) string
	Lookup    func(id string) (*tktypes.HotelHit, error)
}

// Read resolves a hotel by id or URL.
func (r *LDReader) Read(idOrURL string) (*tktypes.HotelView, error) {
	id := strings.TrimSpace(idOrURL)
	if id == "" {
		return nil, fmt.Errorf("id or url required")
	}
	if r.Lookup != nil {
		if hit, err := r.Lookup(id); err == nil && hit != nil {
			return hitToView(hit, r.Brand), nil
		}
	}
	url := id
	if !strings.HasPrefix(url, "http") {
		if r.URLForID != nil {
			url = r.URLForID(id)
		} else {
			url = tkbase.Absolutize(r.BaseURL, "/"+strings.TrimPrefix(id, "/"))
		}
	}
	html, err := r.FetchHTML(url)
	if err != nil {
		return nil, err
	}
	rows := parse.HotelsFromJSONLD(html, r.BaseURL)
	if len(rows) == 0 {
		return nil, fmt.Errorf("hotel not found at %q", idOrURL)
	}
	h := rows[0]
	viewID := id
	if !strings.HasPrefix(id, "http") {
		viewID = strings.TrimPrefix(strings.TrimSuffix(h.URL, "/"), r.BaseURL)
		parts := strings.Split(strings.Trim(viewID, "/"), "/")
		if len(parts) > 0 {
			viewID = parts[len(parts)-1]
		}
	}
	return &tktypes.HotelView{
		ID: viewID, Name: h.Name, Brand: r.Brand,
		Address: h.Address, Stars: h.Stars, HotelURL: h.URL,
	}, nil
}

func hitToView(h *tktypes.HotelHit, brand string) *tktypes.HotelView {
	b := h.Brand
	if b == "" {
		b = brand
	}
	return &tktypes.HotelView{
		ID: h.ID, Name: h.Name, Brand: b, City: h.City, Country: h.Country,
		Stars: h.Stars, HotelURL: h.HotelURL,
		Price: tktypes.PriceInfo{Price: h.Price, Currency: h.Currency, PerNight: true},
	}
}

// LookupFromSearch finds a hit in a full listing search by id or URL fragment.
func LookupFromSearch(search func(string, int, int) (*tktypes.HotelSearchResult, error), id string) (*tktypes.HotelHit, error) {
	res, err := search("", 1, 500)
	if err != nil {
		return nil, err
	}
	q := strings.ToLower(strings.TrimSpace(id))
	for i := range res.Hotels {
		h := &res.Hotels[i]
		if strings.EqualFold(h.ID, id) ||
			strings.Contains(strings.ToLower(h.HotelURL), q) ||
			strings.Contains(strings.ToLower(h.Name), q) {
			return h, nil
		}
	}
	return nil, fmt.Errorf("hotel %q not in listing", id)
}
