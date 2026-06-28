package parse

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strings"
)

var (
	marriottOverviewLinkRE = regexp.MustCompile(`href="(https?://[^"]+/hotels/[a-z0-9]+-[a-z0-9-]+/overview/?)"`)
	hoxtonHotelRE          = regexp.MustCompile(`href="(https?://(?:www\.)?thehoxton\.com/[^"]+)"`)
	wyndhamPropertyRE      = regexp.MustCompile(`"propertyId"\s*:\s*"([^"]+)"[\s\S]{0,600}?"propertyName"\s*:\s*"([^"\\]+)"`)
	wyndhamHotelLinkRE     = regexp.MustCompile(`href="(/[^"]*?/hotels/[^"?#]+)"`)
	easyHotelLinkRE        = regexp.MustCompile(`href="(https?://(?:www\.)?easyhotel\.com/hotels/[^"]+)"`)
	bbHotelLinkRE          = regexp.MustCompile(`href="(/en/(?:gb|fr|de)/[^"]+)"`)
	numaPropertyRE         = regexp.MustCompile(`href="(/en/(?:city|properties)/[^"]+)"`)
	limehomeLinkRE         = regexp.MustCompile(`href="(/en/(?:destinations|locations)/[^"]+)"`)
)

func HotelsFromMarriottOverviewLinks(html, baseURL string) []HotelLD {
	out := HotelsFromMarriottSearch(html, baseURL)
	if len(out) > 0 {
		return out
	}
	host := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	host = strings.TrimSuffix(host, "/")
	re := regexp.MustCompile(`href="((?:https?://` + regexp.QuoteMeta(host) + `)?/hotels/[a-z0-9][a-z0-9-/]*)"`)
	seen := map[string]bool{}
	for _, m := range re.FindAllStringSubmatch(html, -1) {
		u := absolutize(baseURL, m[1])
		if seen[u] {
			continue
		}
		seen[u] = true
		id := pathID(u)
		out = append(out, HotelLD{ID: id, Name: titleFromSlug(strings.ReplaceAll(id, "-", " ")), URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromJSONLD(html, baseURL)
	}
	return out
}

func HotelsFromHoxtonHome(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range hoxtonHotelRE.FindAllStringSubmatch(html, -1) {
		u := strings.TrimSuffix(m[1], "/")
		if strings.Contains(u, "/press") || strings.Contains(u, "/careers") {
			continue
		}
		slug := pathID(u)
		if slug == "" || seen[u] {
			continue
		}
		seen[u] = true
		out = append(out, HotelLD{ID: slug, Name: "The Hoxton " + titleFromSlug(slug), URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromBrandHomeLinks(html, baseURL, "www.thehoxton.com")
	}
	return out
}

func HotelsFromWyndhamDestination(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range wyndhamPropertyRE.FindAllStringSubmatch(html, -1) {
		id, name := m[1], strings.TrimSpace(m[2])
		if seen[id] || name == "" {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{ID: id, Name: name, URL: absolutize(baseURL, "/hotels/"+id)})
	}
	for _, m := range wyndhamHotelLinkRE.FindAllStringSubmatch(html, -1) {
		if strings.Contains(m[1], "destination") {
			continue
		}
		u := absolutize(baseURL, m[1])
		id := pathID(u)
		if seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, HotelLD{ID: id, Name: titleFromSlug(id), URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromJSONLD(html, baseURL)
	}
	return out
}

func HotelsFromBrandHomeLinks(html, baseURL, host string) []HotelLD {
	host = strings.TrimPrefix(strings.TrimPrefix(host, "https://"), "http://")
	host = strings.TrimSuffix(host, "/")
	seen := map[string]bool{}
	var out []HotelLD
	re := regexp.MustCompile(`href="((?:https?://` + regexp.QuoteMeta(host) + `)?/[^"#?]+)"`)
	for _, m := range re.FindAllStringSubmatch(html, -1) {
		u := absolutize(baseURL, m[1])
		low := strings.ToLower(u)
		if !strings.Contains(low, "/hotel") {
			continue
		}
		if seen[u] {
			continue
		}
		seen[u] = true
		id := pathID(u)
		out = append(out, HotelLD{ID: id, Name: titleFromSlug(id), URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromJSONLD(html, baseURL)
	}
	return out
}

func HotelsFromEasyHotelListing(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range easyHotelLinkRE.FindAllStringSubmatch(html, -1) {
		u := strings.TrimSuffix(m[1], "/")
		if seen[u] {
			continue
		}
		seen[u] = true
		id := pathID(u)
		out = append(out, HotelLD{ID: id, Name: titleFromSlug(id), URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromBrandHomeLinks(html, baseURL, "www.easyhotel.com")
	}
	return out
}

func HotelsFromBBHotelsCity(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range bbHotelLinkRE.FindAllStringSubmatch(html, -1) {
		u := absolutize(baseURL, m[1])
		if seen[u] {
			continue
		}
		seen[u] = true
		id := pathID(u)
		out = append(out, HotelLD{ID: id, Name: titleFromSlug(id), URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromJSONLD(html, baseURL)
	}
	return out
}

func HotelsFromNumaCity(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range numaPropertyRE.FindAllStringSubmatch(html, -1) {
		u := absolutize(baseURL, m[1])
		if seen[u] {
			continue
		}
		seen[u] = true
		id := pathID(u)
		out = append(out, HotelLD{ID: id, Name: titleFromSlug(id), URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromJSONLD(html, baseURL)
	}
	return out
}

func HotelsFromLimehomeDestination(html, baseURL string) []HotelLD {
	seen := map[string]bool{}
	var out []HotelLD
	for _, m := range limehomeLinkRE.FindAllStringSubmatch(html, -1) {
		u := absolutize(baseURL, m[1])
		if seen[u] {
			continue
		}
		seen[u] = true
		id := pathID(u)
		out = append(out, HotelLD{ID: id, Name: titleFromSlug(id), URL: u})
	}
	if len(out) == 0 {
		out = HotelsFromJSONLD(html, baseURL)
	}
	return out
}

func HotelsFromSonderGraphQL(body []byte, baseURL string) []HotelLD {
	var root struct {
		Data struct {
			BuildingSearch struct {
				Buildings []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Slug string `json:"slug"`
					City string `json:"city"`
				} `json:"buildings"`
			} `json:"buildingSearch"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &root); err != nil {
		return nil
	}
	out := make([]HotelLD, 0, len(root.Data.BuildingSearch.Buildings))
	for _, b := range root.Data.BuildingSearch.Buildings {
		slug := b.Slug
		if slug == "" {
			slug = b.ID
		}
		u := strings.TrimRight(baseURL, "/") + "/destinations/" + slug
		out = append(out, HotelLD{ID: slug, Name: b.Name, URL: u, Address: b.City})
	}
	return out
}

func HotelsFromSonderHome(html, baseURL string) []HotelLD {
	return HotelsFromBrandHomeLinks(html, baseURL, "www.sonder.com")
}

func BestWesternSearchURL(base, place string) string {
	q := url.QueryEscape(strings.TrimSpace(place))
	return strings.TrimRight(base, "/") + "/en_US/book/hotel-search.html?place=" + q
}

func pathID(u string) string {
	u = strings.TrimSuffix(u, "/")
	parts := strings.Split(u, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
