package parse

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	sercotelTitleRE = regexp.MustCompile(`\\"title\\":\\"([^\\]+)\\"`)
	sercotelCityRE  = regexp.MustCompile(`\\"city\\":\\"([^\\]+)\\"`)
	sercotelStateRE = regexp.MustCompile(`\\"state\\":\\"([^\\]+)\\"`)
	sercotelStarsRE = regexp.MustCompile(`\\"rankingStars\\":\\"([^\\]+)\\"`)
)

// HotelsFromSercotelRSC extracts hotel rows from Magnolia data in Next.js RSC payloads.
func HotelsFromSercotelRSC(html, baseURL string) []HotelLD {
	const marker = `\"@name\":\"sercotel-`
	parts := strings.Split(html, marker)
	if len(parts) < 2 {
		return nil
	}
	seen := map[string]bool{}
	var out []HotelLD
	for _, part := range parts[1:] {
		suffix := part
		if i := strings.Index(suffix, `\"`); i >= 0 {
			suffix = suffix[:i]
		}
		slug := "sercotel-" + suffix
		if seen[slug] {
			continue
		}
		chunk := part
		if len(chunk) > 5000 {
			chunk = chunk[:5000]
		}
		title := sercotelField(sercotelTitleRE, chunk)
		if title == "" || !strings.HasPrefix(title, "Sercotel ") {
			continue
		}
		if strings.Contains(title, "empresas") || strings.Contains(title, "agencias") || strings.Contains(title, "Rewards") {
			continue
		}
		city := sercotelField(sercotelCityRE, chunk)
		if city == "" {
			city = sercotelField(sercotelStateRE, chunk)
		}
		seen[slug] = true
		stars, _ := strconv.ParseFloat(sercotelField(sercotelStarsRE, chunk), 64)
		out = append(out, HotelLD{
			ID:      slug,
			Name:    title,
			Address: city,
			Stars:   stars,
			URL:     absolutize(baseURL, "/es/hoteles-y-destinos"),
		})
	}
	return out
}

func sercotelField(re *regexp.Regexp, chunk string) string {
	if m := re.FindStringSubmatch(chunk); len(m) > 1 {
		return strings.TrimSpace(m[1])
	}
	return ""
}
