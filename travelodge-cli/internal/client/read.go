package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/parse"
)

var jsonLDBlockRE = regexp.MustCompile(`(?s)<script[^>]*type="application/ld\+json"[^>]*>(.*?)</script>`)

type hotelJSONLD struct {
	Type        string `json:"@type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     any    `json:"address"`
	URL         string `json:"url"`
}

// Read returns hotel detail from a hotel page URL or /hotels/... path.
func (c *Client) Read(idOrURL string) (*HotelView, error) {
	raw := strings.TrimSpace(idOrURL)
	if raw == "" {
		return nil, fmt.Errorf("id or url required")
	}

	pageURL, err := c.resolveHotelPageURL(raw)
	if err != nil {
		return nil, err
	}

	html, err := c.FetchHTML(pageURL)
	if err != nil {
		return nil, err
	}

	ld, err := parseHotelJSONLD(html)
	if err != nil {
		return nil, err
	}

	id := raw
	if n, err := strconv.Atoi(raw); err == nil {
		id = strconv.Itoa(n)
	} else if u, err := url.Parse(pageURL); err == nil {
		parts := strings.Split(strings.Trim(u.Path, "/"), "/")
		if len(parts) >= 2 && parts[0] == "hotels" {
			id = parts[1]
		}
	}

	address := formatAddress(ld.Address)
	return &HotelView{
		ID:          id,
		Name:        ld.Name,
		Brand:       brandOrDefault(c.Brand),
		Description: ld.Description,
		Address:     address,
		HotelURL:    pageURL,
	}, nil
}

func (c *Client) resolveHotelPageURL(idOrURL string) (string, error) {
	raw := strings.TrimSpace(idOrURL)
	switch {
	case strings.HasPrefix(raw, "http://") || strings.HasPrefix(raw, "https://"):
		return raw, nil
	case strings.HasPrefix(raw, "/hotels/"):
		return tkbase.Absolutize(c.BaseURL, raw), nil
	case strings.Contains(raw, "/"):
		return tkbase.Absolutize(c.BaseURL, "/hotels/"+strings.TrimPrefix(raw, "/")), nil
	default:
		if _, err := strconv.Atoi(raw); err == nil {
			return "", fmt.Errorf("numeric hotel id %q requires hotel_url from search (e.g. /hotels/318/London-Covent-Garden-hotel)", raw)
		}
		return tkbase.Absolutize(c.BaseURL, "/hotels/"+raw), nil
	}
}

func parseHotelJSONLD(html string) (*hotelJSONLD, error) {
	blocks := jsonLDBlockRE.FindAllStringSubmatch(html, -1)
	for _, m := range blocks {
		if len(m) < 2 {
			continue
		}
		var ld hotelJSONLD
		if err := json.Unmarshal([]byte(m[1]), &ld); err != nil {
			continue
		}
		if strings.EqualFold(ld.Type, "Hotel") && ld.Name != "" {
			return &ld, nil
		}
	}

	rows := parse.HotelsFromJSONLD(html, "")
	if len(rows) == 0 {
		return nil, fmt.Errorf("hotel JSON-LD not found on page")
	}
	return &hotelJSONLD{
		Type: "Hotel",
		Name: rows[0].Name,
		URL:  rows[0].URL,
	}, nil
}

func formatAddress(v any) string {
	switch a := v.(type) {
	case string:
		return a
	case map[string]any:
		parts := []string{
			strVal(a["streetAddress"]),
			strVal(a["addressLocality"]),
			strVal(a["postalCode"]),
			strVal(a["addressCountry"]),
		}
		var out []string
		for _, p := range parts {
			if p != "" {
				out = append(out, p)
			}
		}
		return strings.Join(out, ", ")
	default:
		return ""
	}
}

func strVal(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
