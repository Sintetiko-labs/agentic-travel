package parse

import "testing"

func TestHotelsFromIHGSearchMnemonic(t *testing.T) {
	html := `{"hotelMnemonic":"MADCP","name":"Holiday Inn Madrid - Calle Alcala"}`
	rows := HotelsFromIHGSearch(html, "https://www.ihg.com")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].ID != "madcp" {
		t.Fatalf("id=%q", rows[0].ID)
	}
	if rows[0].Name != "Holiday Inn Madrid - Calle Alcala" {
		t.Fatalf("name=%q", rows[0].Name)
	}
}
