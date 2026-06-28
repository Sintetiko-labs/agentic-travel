package parse

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var ngStateRE = regexp.MustCompile(`<script id="ng-state"[^>]*>(.*?)</script>`)

// RIUHotel is a hotel row parsed from RIU Angular ng-state.
type RIUHotel struct {
	ID          string
	Name        string
	Slug        string
	City        string
	Country     string
	Stars       float64
	Price       float64
	Currency    string
	Rating      float64
	ReviewCount int
	ImageURL    string
}

// ParseRIUNgState extracts hotel rows from a RIU destination page.
func ParseRIUNgState(html, baseURL, query string) []RIUHotel {
	m := ngStateRE.FindStringSubmatch(html)
	if m == nil {
		return nil
	}
	var state map[string]json.RawMessage
	if err := json.Unmarshal([]byte(m[1]), &state); err != nil {
		return nil
	}
	q := strings.ToLower(strings.TrimSpace(query))
	var hotels []RIUHotel
	var filterRows []RIUHotel
	for key, raw := range state {
		if !strings.HasPrefix(key, "KEY_HOTELS") && !strings.Contains(key, "hotels?") && !strings.HasPrefix(key, "KEY_HOTEL") {
			continue
		}
		rows := unmarshalRIURows(raw)
		if strings.HasPrefix(key, "KEY_HOTELS_FILTER") {
			filterRows = append(filterRows, mapRIUHotels(rows, baseURL, q)...)
			continue
		}
		hotels = append(hotels, mapRIUHotels(rows, baseURL, q)...)
	}
	if len(filterRows) > 0 {
		return dedupeRIU(filterRows)
	}
	return dedupeRIU(hotels)
}

func unmarshalRIURows(raw json.RawMessage) []map[string]any {
	var list []map[string]any
	if err := json.Unmarshal(raw, &list); err == nil && len(list) > 0 {
		return list
	}
	var wrapped struct {
		Body []map[string]any `json:"body"`
	}
	if err := json.Unmarshal(raw, &wrapped); err == nil && len(wrapped.Body) > 0 {
		return wrapped.Body
	}
	return nil
}

func mapRIUHotels(rows []map[string]any, baseURL, query string) []RIUHotel {
	out := make([]RIUHotel, 0, len(rows))
	for _, row := range rows {
		name := firstString(row, "name", "hotelName")
		if name == "" {
			continue
		}
		city := nestedString(row, "destination", "name")
		if city == "" {
			city = str(row["hotelDestination"])
		}
		country := nestedString(row, "country", "name")
		if country == "" {
			country = str(row["hotelCountry"])
		}
		if query != "" && !strings.Contains(strings.ToLower(name), query) &&
			!strings.Contains(strings.ToLower(city), query) &&
			!strings.Contains(strings.ToLower(country), query) {
			continue
		}
		h := RIUHotel{
			ID:       firstString(row, "id"),
			Name:     name,
			Slug:     firstString(row, "slug", "hotelSlug"),
			City:     city,
			Country:  country,
			Stars:    num(row["stars"]),
			Currency: "EUR",
		}
		if p, ok := row["price"].(map[string]any); ok {
			h.Price = num(p["amount"])
			if c := str(p["currency"]); c != "" {
				h.Currency = c
			}
		}
		if ta, ok := row["trip_advisor"].(map[string]any); ok {
			h.Rating = num(ta["rating"])
			h.ReviewCount = int(num(ta["reviews"]))
		}
		if gal, ok := row["gallery"].([]any); ok && len(gal) > 0 {
			if im, ok := gal[0].(map[string]any); ok {
				if img, ok := im["image"].(map[string]any); ok {
					h.ImageURL = absolutize(baseURL, str(img["path"]))
				}
			}
		}
		if h.ImageURL == "" {
			if hi, ok := row["highlights_image"].(map[string]any); ok {
				h.ImageURL = absolutize(baseURL, str(hi["path"]))
			} else {
				h.ImageURL = absolutize(baseURL, str(row["highlightsImage"]))
			}
		}
		out = append(out, h)
	}
	return out
}

func dedupeRIU(in []RIUHotel) []RIUHotel {
	seen := map[string]bool{}
	out := make([]RIUHotel, 0, len(in))
	for _, h := range in {
		key := h.ID
		if key == "" {
			key = h.Slug
		}
		if key == "" || seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, h)
	}
	return out
}

func firstString(m map[string]any, keys ...string) string {
	for _, k := range keys {
		if v := str(m[k]); v != "" {
			return v
		}
	}
	return ""
}

func nestedString(m map[string]any, key, sub string) string {
	if v, ok := m[key].(map[string]any); ok {
		return str(v[sub])
	}
	return ""
}

func str(v any) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	case float64:
		return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", t), "0"), ".")
	default:
		b, _ := json.Marshal(t)
		if len(b) >= 2 && b[0] == '"' {
			var s string
			_ = json.Unmarshal(b, &s)
			return s
		}
		return string(b)
	}
}

func num(v any) float64 {
	switch t := v.(type) {
	case float64:
		return t
	case string:
		var f float64
		_ = json.Unmarshal([]byte(t), &f)
		return f
	default:
		return 0
	}
}
