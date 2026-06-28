package parse

import (
	"encoding/json"
	"strings"
)

// H10Hotel is a hotel row parsed from H10 Angular ng-state menu-es.
type H10Hotel struct {
	ID       string
	Name     string
	Slug     string
	City     string
	Country  string
	Stars    float64
	Price    float64
	Currency string
	ImageURL string
}

// ParseH10NgState extracts hotels from H10 destination pages (menu-es key).
func ParseH10NgState(html, baseURL, query string) []H10Hotel {
	m := ngStateRE.FindStringSubmatch(html)
	if m == nil {
		return nil
	}
	var state map[string]json.RawMessage
	if err := json.Unmarshal([]byte(m[1]), &state); err != nil {
		return nil
	}
	raw, ok := state["menu-es"]
	if !ok {
		for k, v := range state {
			if strings.HasSuffix(k, "menu-es") {
				raw = v
				ok = true
				break
			}
		}
	}
	if !ok {
		return nil
	}
	var menu struct {
		DestinationCategories []struct {
			Name         string `json:"name"`
			Destinations []struct {
				Name   string           `json:"name"`
				Hotels []map[string]any `json:"hotels"`
			} `json:"destinations"`
		} `json:"destinationCategories"`
	}
	if err := json.Unmarshal(raw, &menu); err != nil {
		return nil
	}
	q := strings.ToLower(strings.TrimSpace(query))
	var out []H10Hotel
	for _, cat := range menu.DestinationCategories {
		country := cat.Name
		for _, dest := range cat.Destinations {
			city := dest.Name
			if q != "" && !strings.Contains(strings.ToLower(city), q) &&
				!strings.Contains(strings.ToLower(country), q) {
				continue
			}
			for _, row := range dest.Hotels {
				name := firstString(row, "name")
				if name == "" {
					continue
				}
				if q != "" && !strings.Contains(strings.ToLower(name), q) &&
					!strings.Contains(strings.ToLower(city), q) &&
					!strings.Contains(strings.ToLower(country), q) {
					continue
				}
				h := H10Hotel{
					ID:      firstString(row, "id"),
					Name:    name,
					Slug:    firstString(row, "nameUrl"),
					City:    firstString(row, "city", "destination"),
					Country: country,
					Stars:   num(row["category"]),
				}
				if h.City == "" {
					h.City = city
				}
				if img, ok := row["summaryImage"].(map[string]any); ok {
					h.ImageURL = absolutize(baseURL, str(img["url"]))
				}
				if webs, ok := row["webs"].([]any); ok && len(webs) > 0 {
					if w0, ok := webs[0].(map[string]any); ok {
						if mp := num(w0["minPrice"]); mp > 0 {
							h.Price = mp
							h.Currency = "EUR"
						}
					}
				}
				out = append(out, h)
			}
		}
	}
	return dedupeH10(out)
}

func dedupeH10(in []H10Hotel) []H10Hotel {
	seen := map[string]bool{}
	out := make([]H10Hotel, 0, len(in))
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
