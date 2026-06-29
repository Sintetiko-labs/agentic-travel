package parse

import (
	"regexp"
	"strings"
)

var (
	ihgMnemonicJSONRE = regexp.MustCompile(`"hotelMnemonic"\s*:\s*"([A-Z0-9]{3,8})"[\s\S]{0,800}?"name"\s*:\s*"([^"\\]+)"`)
	ihgAltMnemonicRE  = regexp.MustCompile(`"name"\s*:\s*"([^"\\]+)"[\s\S]{0,800}?"hotelMnemonic"\s*:\s*"([A-Z0-9]{3,8})"`)
	ihgHotelCodeRE    = regexp.MustCompile(`"hotelCode"\s*:\s*"([A-Z0-9]{3,8})"[\s\S]{0,800}?"propertyName"\s*:\s*"([^"\\]+)"`)
	ihgHotelLinkRE    = regexp.MustCompile(`href="(/hotels/[^"]+/hotels/[^"]+/hotel/details\?[^"]*hotelCode=([A-Z0-9]+)[^"]*)"`)
)

// HotelsFromIHGSearch extracts hotel rows from IHG find-hotels list HTML or embedded JSON.
func HotelsFromIHGSearch(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range ihgMnemonicJSONRE.FindAllStringSubmatch(html, -1) {
		id, name := m[1], strings.TrimSpace(m[2])
		if seen[id] || name == "" {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{
			ID:   strings.ToLower(id),
			Name: name,
			URL:  ihgHotelURL(baseURL, id),
		})
	}
	for _, m := range ihgAltMnemonicRE.FindAllStringSubmatch(html, -1) {
		name, id := strings.TrimSpace(m[1]), m[2]
		if seen[id] || name == "" {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{
			ID:   strings.ToLower(id),
			Name: name,
			URL:  ihgHotelURL(baseURL, id),
		})
	}
	for _, m := range ihgHotelCodeRE.FindAllStringSubmatch(html, -1) {
		id, name := m[1], strings.TrimSpace(m[2])
		if seen[id] || name == "" {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{
			ID:   strings.ToLower(id),
			Name: name,
			URL:  ihgHotelURL(baseURL, id),
		})
	}
	for _, m := range ihgHotelLinkRE.FindAllStringSubmatch(html, -1) {
		path, id := m[1], m[2]
		if seen[id] {
			continue
		}
		seen[id] = true
		u := absolutize(baseURL, path)
		name := titleFromSlug(strings.ToLower(id))
		out = append(out, HotelLD{ID: strings.ToLower(id), Name: name, URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromMarriottSearch(html, baseURL)
	}
	return out
}

func ihgHotelURL(base, code string) string {
	code = strings.ToUpper(strings.TrimSpace(code))
	return strings.TrimRight(base, "/") + "/hotels/gb/en/find-hotels/hotel/details?hotelCode=" + code
}
