package parse

import "testing"

func TestHotelsFromMarriottSearchJSON(t *testing.T) {
	html := `{"propertyCode":"LONCY","name":"London Marriott Hotel County Hall"}`
	rows := HotelsFromMarriottSearch(html, "https://www.marriott.com")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].ID != "loncy" {
		t.Fatalf("id=%q", rows[0].ID)
	}
	if rows[0].Name != "London Marriott Hotel County Hall" {
		t.Fatalf("name=%q", rows[0].Name)
	}
}
