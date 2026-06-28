package parse

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var hotelNameRE = regexp.MustCompile(`"@type"\s*:\s*"Hotel"\s*,\s*"name"\s*:\s*"([^"\\]*(?:\\.[^"\\]*)*)"`)

// HotelLD is a minimal JSON-LD hotel row extracted from HTML.
type HotelLD struct {
	Name       string
	Stars      float64
	Address    string
	URL        string
	ImageURL   string
}

// HotelsFromJSONLD scans HTML for embedded Hotel JSON-LD name fields and nearby metadata.
func HotelsFromJSONLD(html, baseURL string) []HotelLD {
	names := hotelNameRE.FindAllStringSubmatch(html, -1)
	if len(names) == 0 {
		return nil
	}
	seen := map[string]bool{}
	out := make([]HotelLD, 0, len(names))
	for _, m := range names {
		name := strings.ReplaceAll(m[1], `\"`, `"`)
		if seen[name] {
			continue
		}
		seen[name] = true
		h := HotelLD{Name: name}
		chunk := extractHotelChunk(html, m[0])
		h.Stars = jsonLDStars(chunk)
		h.Address = jsonLDString(chunk, "address")
		h.URL = absolutize(baseURL, jsonLDString(chunk, "url"))
		if imgs := jsonLDImages(chunk); len(imgs) > 0 {
			h.ImageURL = absolutize(baseURL, imgs[0])
		}
		out = append(out, h)
	}
	return out
}

func extractHotelChunk(html, anchor string) string {
	idx := strings.Index(html, anchor)
	if idx < 0 {
		return anchor
	}
	start := idx
	if start > 400 {
		start -= 400
	}
	end := idx + len(anchor) + 1200
	if end > len(html) {
		end = len(html)
	}
	return html[start:end]
}

func jsonLDStars(chunk string) float64 {
	var wrapper struct {
		StarRating json.RawMessage `json:"starRating"`
	}
	// try object form
	var obj struct {
		StarRating float64 `json:"starRating"`
	}
	if err := json.Unmarshal([]byte(chunk), &obj); err == nil && obj.StarRating > 0 {
		return obj.StarRating
	}
	_ = json.Unmarshal(wrapper.StarRating, &obj.StarRating)
	if re := regexp.MustCompile(`"starRating"\s*:\s*([0-9]+(?:\.[0-9]+)?)`); re != nil {
		if m := re.FindStringSubmatch(chunk); len(m) > 1 {
			var f float64
			_, _ = fmt.Sscanf(m[1], "%f", &f)
			if f > 0 {
				return f
			}
		}
	}
	return 0
}

func jsonLDString(chunk, key string) string {
	re := regexp.MustCompile(`"` + key + `"\s*:\s*"([^"\\]*(?:\\.[^"\\]*)*)"`)
	if m := re.FindStringSubmatch(chunk); len(m) > 1 {
		return strings.ReplaceAll(m[1], `\"`, `"`)
	}
	return ""
}

func jsonLDImages(chunk string) []string {
	re := regexp.MustCompile(`"images"\s*:\s*\[(.*?)\]`)
	m := re.FindStringSubmatch(chunk)
	if len(m) < 2 {
		return nil
	}
	imgRE := regexp.MustCompile(`"([^"]+)"`)
	var imgs []string
	for _, im := range imgRE.FindAllStringSubmatch(m[1], -1) {
		if len(im) > 1 {
			imgs = append(imgs, im[1])
		}
	}
	return imgs
}

func absolutize(base, path string) string {
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	if strings.HasPrefix(path, "//") {
		return "https:" + path
	}
	if strings.HasPrefix(path, "/") {
		return strings.TrimRight(base, "/") + path
	}
	return strings.TrimRight(base, "/") + "/" + path
}
