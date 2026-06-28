package hotel

import (
	"strings"

	"github.com/fbelchi/travelkit/parse"
)

// FilterByBrand keeps rows whose name contains the brand token.
func FilterByBrand(rows []parse.HotelLD, brand string) []parse.HotelLD {
	b := strings.TrimSpace(brand)
	if b == "" {
		return rows
	}
	low := strings.ToLower(b)
	out := make([]parse.HotelLD, 0, len(rows))
	for _, h := range rows {
		if strings.Contains(strings.ToLower(h.Name), low) {
			out = append(out, h)
		}
	}
	return out
}
