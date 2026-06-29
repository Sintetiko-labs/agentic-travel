package parse
import (
	"encoding/json"
	"regexp"
	"strings"
)

var accorHotelCodeRE = regexp.MustCompile(`accorhotels/([a-z]{3}_[a-z]_[0-9]+)`)

// HotelsFromAccorSearch parses Accor search JSON or falls back to destination HTML/JSON-LD.
func HotelsFromAccorSearch(html, baseURL string) []HotelLD {
	html = strings.TrimSpace(html)
	if strings.HasPrefix(html, "{") {
		var row struct {
			HotelCode string `json:"hotelCode"`
			Name      string `json:"name"`
		}
		if json.Unmarshal([]byte(html), &row) == nil && row.HotelCode != "" {
			return []HotelLD{{
				ID:   row.HotelCode,
				Name: row.Name,
				URL:  absolutize(baseURL, "/hotels/"+row.HotelCode),
			}}
		}
	}
	if out := HotelsFromJSONLD(html, baseURL); len(out) > 0 {
		return out
	}
	return HotelsFromAccorDestination(html, baseURL)
}

func HotelsFromAccorSSR(html, baseURL string) []HotelLD {
 seen := map[string]bool{}; var out []HotelLD
 for _, m := range accorHotelCodeRE.FindAllStringSubmatch(html, -1) {
  code := m[1]; if seen[code] { continue }; seen[code]=true
  out = append(out, HotelLD{ID: code, Name: titleFromSlug(strings.ReplaceAll(code,"_"," ")), URL: absolutize(baseURL,"/hotels/"+code)})
 }
 return out
}
func HotelsFromAccorDestination(html, baseURL string) []HotelLD {
 if out := HotelsFromJSONLD(html, baseURL); len(out)>0 { return out }
 if out := HotelsFromAccorSSR(html, baseURL); len(out)>0 { return out }
 return HotelsFromBrandHomeLinks(html, baseURL, "all.accor.com")
}
