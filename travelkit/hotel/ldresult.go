package hotel

import (
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/parse"
	tktypes "github.com/fbelchi/travelkit/types"
)

// FilterByBrand keeps rows whose name contains the brand token.
func FilterByBrand(rows []parse.HotelLD, brand string) []parse.HotelLD {
	b := strings.TrimSpace(brand)
	if b == "" {
		return rows
	}
	low := strings.ToLower(b)
	out := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if strings.Contains(strings.ToLower(h.Name), low) {
			out = append(out, h)
		}
	}
	return out
}

// FilterHotelLD keeps rows matching destination query (name, address, url).
func FilterHotelLD(rows []parse.HotelLD, query string) []parse.HotelLD {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return rows
	}
	terms := strings.FieldsFunc(q, func(r rune) bool { return r == ',' || r == '|' || r == '&' })
	if len(terms) == 0 {
		terms = []string{q}
	}
	out := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		blob := strings.ToLower(strings.Join([]string{h.Name, h.Address, h.URL, h.ID}, " "))
		for _, t := range terms {
			t = strings.TrimSpace(t)
			if t != "" && strings.Contains(blob, t) {
				out = append(out, h)
				break
			}
		}
	}
	return out
}

// LDToResult paginates HotelLD rows into a normalized search result.
// When brand is empty, parentSlug selects sub-brand inference from hotel names.
func LDToResult(rows []parse.HotelLD, query string, page, pageSize int, brand, baseURL, source string) *tktypes.HotelSearchResult {
	return LDToResultParent(rows, query, page, pageSize, brand, "", baseURL, source)
}

// LDToResultParent is like LDToResult but infers sub-brands when brand is empty.
func LDToResultParent(rows []parse.HotelLD, query string, page, pageSize int, brand, parentSlug, baseURL, source string) *tktypes.HotelSearchResult {
	total := len(rows)
	start := (page - 1) * pageSize
	if start > total {
		start = total
	}
	end := start + pageSize
	if end > total {
		end = total
	}
	hits := make([]tktypes.HotelHit, 0, end-start)
	for _, h := range rows[start:end] {
		id := h.ID
		if id == "" {
			id = slugFromURL(h.URL)
		}
		b := brand
		if b == "" && parentSlug != "" {
			b = InferBrand(parentSlug, h.Name)
		}
		if b == "" {
			b = h.Name
		}
		hits = append(hits, tktypes.HotelHit{
			ID:       id,
			Name:     h.Name,
			Brand:    b,
			City:     query,
			Stars:    h.Stars,
			HotelURL: tkbase.Absolutize(baseURL, h.URL),
			ImageURL: h.ImageURL,
		})
	}
	return &tktypes.HotelSearchResult{
		Query:    query,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		HasNext:  total > page*pageSize,
		Hotels:   hits,
		Brand:    brand,
		Source:   source,
	}
}

func slugFromURL(raw string) string {
	if raw == "" {
		return ""
	}
	parts := strings.Split(strings.Trim(raw, "/"), "/")
	if len(parts) == 0 {
		return raw
	}
	return parts[len(parts)-1]
}
