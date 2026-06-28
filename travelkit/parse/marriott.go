package parse

import (
	"regexp"
	"strings"
)

var (
	marriottPropertyJSONRE = regexp.MustCompile(`"propertyCode"\s*:\s*"([A-Z0-9]{3,8})"[\s\S]{0,800}?"name"\s*:\s*"([^"\\]+)"`)
	marriottHotelLinkRE    = regexp.MustCompile(`href="(https://www\.marriott\.com/[^"]*/hotels/([a-z0-9]+)-[^"/]+/overview/?)"`)
	marriottAltJSONRE      = regexp.MustCompile(`"hotelName"\s*:\s*"([^"\\]+)"[\s\S]{0,400}?"propertyCode"\s*:\s*"([A-Z0-9]{3,8})"`)
)

// HotelsFromMarriottSearch extracts hotel rows from findHotels HTML or embedded JSON.
func HotelsFromMarriottSearch(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range marriottPropertyJSONRE.FindAllStringSubmatch(html, -1) {
		id, name := m[1], strings.TrimSpace(m[2])
		if seen[id] || name == "" {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{
			ID:   strings.ToLower(id),
			Name: name,
			URL:  marriottHotelURL(baseURL, id, name),
		})
	}
	for _, m := range marriottAltJSONRE.FindAllStringSubmatch(html, -1) {
		name, id := strings.TrimSpace(m[1]), m[2]
		if seen[id] || name == "" {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{
			ID:   strings.ToLower(id),
			Name: name,
			URL:  marriottHotelURL(baseURL, id, name),
		})
	}
	for _, m := range marriottHotelLinkRE.FindAllStringSubmatch(html, -1) {
		url, code := m[1], strings.ToLower(m[2])
		if seen[code] {
			continue
		}
		seen[code] = true
		slug := strings.TrimPrefix(strings.TrimSuffix(url, "/overview"), "/overview/")
		parts := strings.Split(slug, "/")
		name := titleFromSlug(strings.TrimPrefix(parts[len(parts)-1], code+"-"))
		out = append(out, HotelLD{
			ID:   code,
			Name: name,
			URL:  url,
		})
	}
	return out
}

func marriottHotelURL(base, code, name string) string {
	slug := strings.ToLower(code) + "-" + strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	slug = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(slug, "")
	return strings.TrimRight(base, "/") + "/en-us/hotels/" + slug + "/overview/"
}
