package destination

import "strings"

// Aliases maps a lowercase city token to alternate spellings/slugs used on hotel sites.
var Aliases = map[string][]string{
	"palma":        {"mallorca", "majorica", "palma-de-mallorca", "islas-baleares", "maspalomas"},
	"barcelona":    {"bcn", "cataluna", "catalunya"},
	"valencia":     {"valència"},
	"madrid":       {"comunidad-de-madrid"},
	"gran canaria": {"gran-canaria", "maspalomas", "meloneras"},
}

// Expand returns query plus known aliases (deduped, lowercase slug form included).
func Expand(query string) []string {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil
	}
	seen := map[string]bool{strings.ToLower(q): true}
	out := []string{q}
	key := strings.ToLower(q)
	if al, ok := Aliases[key]; ok {
		for _, a := range al {
			if !seen[a] {
				seen[a] = true
				out = append(out, a)
			}
		}
	}
	slug := strings.ToLower(strings.ReplaceAll(q, " ", "-"))
	if !seen[slug] {
		out = append(out, slug)
	}
	return out
}

// MatchQuery reports whether a hotel field matches the destination query or aliases.
func MatchQuery(query, name, url, address, id string) bool {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return true
	}
	blob := strings.ToLower(strings.Join([]string{name, url, address, id}, " "))
	if strings.Contains(blob, q) {
		return true
	}
	for _, alt := range Aliases[q] {
		if strings.Contains(blob, alt) {
			return true
		}
	}
	return false
}
