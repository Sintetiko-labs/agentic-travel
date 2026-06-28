package parse

import "testing"

func TestHotelsFromHiltonLocations(t *testing.T) {
	html := `<a href="https://www.hilton.com/en/hotels/lonnmnd-nomad-london/">x<div data-testid="listViewPropertyName">NoMad London</div></a>`
	rows := HotelsFromHiltonLocations(html, "https://www.hilton.com")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].Name != "NoMad London" {
		t.Fatalf("name=%q", rows[0].Name)
	}
	if rows[0].ID != "lonnmnd-nomad-london" {
		t.Fatalf("id=%q", rows[0].ID)
	}
}
