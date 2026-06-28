package parse

import (
	"regexp"
	"strings"
)

var (
	silkenHotelRE = regexp.MustCompile(`data-hotel="(\d+)"\s+data-id="([^"]+)"`)
	silkenSpanRE  = regexp.MustCompile(`<span[^>]*>([^<]{4,80})</span>`)
)

// HotelsFromSilkenCards extracts hotel rows from Silken listing data attributes.
func HotelsFromSilkenCards(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range silkenHotelRE.FindAllStringSubmatch(html, -1) {
		id, slug := m[1], m[2]
		if seen[id] {
			continue
		}
		seen[id] = true
		name := silkenNameNear(html, slug)
		if name == "" {
			name = titleFromSlug(strings.TrimPrefix(slug, "slk_"))
		}
		out = append(out, HotelLD{
			ID:   id,
			Name: name,
			URL:  absolutize(baseURL, "/es/hoteles"),
		})
	}
	return out
}

func silkenNameNear(html, slug string) string {
	anchor := `data-id="` + slug + `"`
	idx := strings.Index(html, anchor)
	if idx < 0 {
		return ""
	}
	end := idx + 900
	if end > len(html) {
		end = len(html)
	}
	chunk := html[idx:end]
	for _, m := range silkenSpanRE.FindAllStringSubmatch(chunk, -1) {
		name := strings.TrimSpace(m[1])
		if name == "" || strings.EqualFold(name, "Ver hotel") {
			continue
		}
		return name
	}
	return ""
}
