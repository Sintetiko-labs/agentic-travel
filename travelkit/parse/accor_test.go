package parse

import "testing"

func TestHotelsFromAccorSearchJSON(t *testing.T) {
	html := `{"hotelCode":"A5E7","name":"Novotel Madrid Center"}`
	rows := HotelsFromAccorSearch(html, "https://all.accor.com")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].Name != "Novotel Madrid Center" {
		t.Fatalf("name=%q", rows[0].Name)
	}
}
