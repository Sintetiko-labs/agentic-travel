package hotel

import (
	"testing"

	"github.com/fbelchi/travelkit/parse"
)

func TestInferBrandAccor(t *testing.T) {
	if got := InferBrand("accor", "Novotel Madrid Center"); got != "Novotel" {
		t.Fatalf("got %q", got)
	}
}

func TestInferBrandMarriott(t *testing.T) {
	if got := InferBrand("marriott", "London Marriott Hotel County Hall"); got != "Marriott" {
		t.Fatalf("got %q", got)
	}
	if got := InferBrand("marriott", "The Ritz-Carlton London"); got != "The Ritz-Carlton" {
		t.Fatalf("got %q", got)
	}
}

func TestFilterHotelLD(t *testing.T) {
	rows := []parse.HotelLD{
		{Name: "Hotel Alpha", Address: "London"},
		{Name: "Hotel Beta", Address: "Paris"},
	}
	out := FilterHotelLD(rows, "London")
	if len(out) != 1 || out[0].Name != "Hotel Alpha" {
		t.Fatalf("filter: %+v", out)
	}
}
