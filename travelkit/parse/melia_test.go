package parse

import "testing"

func TestHotelsFromMeliaDirectory(t *testing.T) {
	html := `<a href="/es/hoteles/espana/madrid/melia-castilla">Meliá Castilla</a>
<a href="/es/hoteles/espana/madrid/innside-madrid-genova">INNSiDE Madrid Génova</a>
<a href="/es/hoteles/espana/barcelona/melia-barcelona-sarria">Meliá Barcelona Sarrià</a>`
	rows := HotelsFromMeliaDirectory(html, "https://www.melia.com", "Madrid")
	if len(rows) != 2 {
		t.Fatalf("got %d rows, want 2: %+v", len(rows), rows)
	}
}
