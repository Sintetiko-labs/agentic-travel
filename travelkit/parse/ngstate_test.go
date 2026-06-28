package parse

import (
	"encoding/json"
	"testing"
)

func TestParseRIUNgStateFilterKey(t *testing.T) {
	state := map[string]json.RawMessage{
		"KEY_HOTELS_FILTER_582": json.RawMessage(`[{"id":"582","name":"Hotel Riu Plaza España","slug":"hotel-riu-plaza-espana","destination":{"name":"Madrid"},"country":{"name":"España"},"price":{"amount":152.15,"currency":"EUR"},"trip_advisor":{"rating":"4.4","reviews":"4492"}}]`),
	}
	b, _ := json.Marshal(state)
	html := `<script id="ng-state" type="application/json">` + string(b) + `</script>`
	rows := ParseRIUNgState(html, "https://www.riu.com", "")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].Name != "Hotel Riu Plaza España" {
		t.Fatalf("name=%q", rows[0].Name)
	}
}
