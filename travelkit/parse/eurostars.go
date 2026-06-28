package parse

import (
	"regexp"
	"strconv"
	"strings"
)

var eurostarsHotelRE = regexp.MustCompile(`\{"id":(\d+),"code":"[^"]+","name":"([^"]+)","stars":"([^"]*)","slug":"([^"]+)"`)

// HotelsFromEurostarsEmbedded extracts hotel rows from homepage embedded JSON.
func HotelsFromEurostarsEmbedded(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range eurostarsHotelRE.FindAllStringSubmatch(html, -1) {
		id, name, starsRaw, slugRaw := m[1], m[2], m[3], m[4]
		name = unescapeJSON(name)
		slugRaw = unescapeJSON(slugRaw)
		url := absolutize(baseURL, slugRaw)
		if url == "" || seen[url] {
			continue
		}
		seen[url] = true
		slug := slugFromHotelURL(url)
		out = append(out, HotelLD{
			ID:    firstNonEmpty(id, slug),
			Name:  name,
			Stars: parseStarString(starsRaw),
			URL:   url,
		})
	}
	return out
}

func unescapeJSON(s string) string {
	s = strings.ReplaceAll(s, `\u00e1`, "á")
	s = strings.ReplaceAll(s, `\u00e9`, "é")
	s = strings.ReplaceAll(s, `\u00ed`, "í")
	s = strings.ReplaceAll(s, `\u00f3`, "ó")
	s = strings.ReplaceAll(s, `\u00fa`, "ú")
	s = strings.ReplaceAll(s, `\u00c1`, "Á")
	s = strings.ReplaceAll(s, `\u00c9`, "É")
	s = strings.ReplaceAll(s, `\u00cd`, "Í")
	s = strings.ReplaceAll(s, `\u00d3`, "Ó")
	s = strings.ReplaceAll(s, `\u00da`, "Ú")
	s = strings.ReplaceAll(s, `\u00f1`, "ñ")
	s = strings.ReplaceAll(s, `\u00d1`, "Ñ")
	s = strings.ReplaceAll(s, `\/`, "/")
	return s
}

func parseStarString(raw string) float64 {
	raw = strings.TrimSuffix(strings.TrimSpace(raw), "*")
	if raw == "" {
		return 0
	}
	f, _ := strconv.ParseFloat(raw, 64)
	return f
}

func slugFromHotelURL(raw string) string {
	raw = strings.TrimSuffix(raw, ".html")
	parts := strings.Split(strings.Trim(raw, "/"), "/")
	if len(parts) == 0 {
		return raw
	}
	return parts[len(parts)-1]
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
