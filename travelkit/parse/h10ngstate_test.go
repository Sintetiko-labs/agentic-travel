package parse

import (
	"encoding/json"
	"testing"
)

func TestParseH10NgStateBarcelona(t *testing.T) {
	menu := map[string]any{
		"destinationCategories": []map[string]any{{
			"name": "España",
			"destinations": []map[string]any{{
				"name": "Barcelona",
				"hotels": []map[string]any{{
					"id": "abc", "name": "H10 Art Gallery", "nameUrl": "h10-art-gallery",
					"city": "Barcelona", "category": 4,
					"summaryImage": map[string]any{"url": "https://pro-static.h10hotels.com/gallery/x.jpg"},
					"webs": []map[string]any{{"minPrice": 40}},
				}},
			}},
		}},
	}
	state := map[string]json.RawMessage{
		"menu-es": mustJSON(menu),
	}
	b, _ := json.Marshal(state)
	html := `<script id="ng-state" type="application/json">` + string(b) + `</script>`
	rows := ParseH10NgState(html, "https://www.h10hotels.com", "Barcelona")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].Name != "H10 Art Gallery" {
		t.Fatalf("name=%q", rows[0].Name)
	}
}

func mustJSON(v any) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}
