package parse

import "testing"

func TestHotelsFromEurostarsEmbedded(t *testing.T) {
	html := `{"id":126,"code":"SYN126","name":"Exe Madrid Norte","stars":"4*","slug":"https:\/\/www.eurostarshotels.com\/exe-madrid-norte.html"}`
	rows := HotelsFromEurostarsEmbedded(html, "https://www.eurostarshotels.com")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].Name != "Exe Madrid Norte" {
		t.Fatalf("name=%q", rows[0].Name)
	}
	if rows[0].Stars != 4 {
		t.Fatalf("stars=%v", rows[0].Stars)
	}
}

func TestHotelsFromVincciLinks(t *testing.T) {
	html := `<a href="/es/hoteles/madrid/vincci-capitol/">Capitol</a>`
	rows := HotelsFromVincciLinks(html, "https://www.vinccihoteles.com")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].ID != "vincci-capitol" {
		t.Fatalf("id=%q", rows[0].ID)
	}
	if rows[0].Address != "madrid" {
		t.Fatalf("city=%q", rows[0].Address)
	}
}

func TestHotelsFromSilkenCards(t *testing.T) {
	html := `data-hotel="1" data-id="slk_ramblas"><span>Ramblas Barcelona</span>`
	rows := HotelsFromSilkenCards(html, "https://www.hoteles-silken.com")
	if len(rows) != 1 || rows[0].Name != "Ramblas Barcelona" {
		t.Fatalf("got %+v", rows)
	}
}

func TestHotelsFromSercotelRSC(t *testing.T) {
	html := `\"@name\":\"sercotel-madrid-aeropuerto\",\"@nodeType\":\"mgnl:hotel\",\"title\":\"Sercotel Madrid Aeropuerto\",\"city\":\"Madrid\",\"rankingStars\":\"4\"`
	rows := HotelsFromSercotelRSC(html, "https://www.sercotelhoteles.com")
	if len(rows) != 1 {
		t.Fatalf("got %d rows", len(rows))
	}
	if rows[0].Name != "Sercotel Madrid Aeropuerto" {
		t.Fatalf("name=%q", rows[0].Name)
	}
}

func TestHotelsFromEsCityHotelLinks(t *testing.T) {
	html := `<a href="/es/hoteles/madrid/petit-palace-mayfair/">Mayfair</a>`
	rows := HotelsFromEsCityHotelLinks(html, "https://www.petitpalace.com")
	if len(rows) != 1 || rows[0].ID != "petit-palace-mayfair" {
		t.Fatalf("got %+v", rows)
	}
}

func TestHotelsFromParadoresLinks(t *testing.T) {
	html := `<a href="/es/parador/segovia">Segovia</a>`
	rows := HotelsFromParadoresLinks(html, "https://www.parador.es")
	if len(rows) != 1 || rows[0].ID != "segovia" {
		t.Fatalf("got %+v", rows)
	}
}

func TestHotelsFromIberostarDirectory(t *testing.T) {
	html := `<a href="/es/hoteles/madrid/iberostar-las-letras">Iberostar Las Letras</a>`
	rows := HotelsFromIberostarDirectory(html, "https://www.iberostar.com", "Madrid")
	if len(rows) != 1 || rows[0].Name != "Iberostar Las Letras" {
		t.Fatalf("got %+v", rows)
	}
}

func TestHotelsFromNHDirectory(t *testing.T) {
	html := `<a href="/es/hotel/nh-collection-madrid-suecia">NH</a>`
	rows := HotelsFromNHDirectory(html, "https://www.nh-hotels.com", "Madrid")
	if len(rows) != 1 || rows[0].ID != "nh-collection-madrid-suecia" {
		t.Fatalf("got %+v", rows)
	}
}
