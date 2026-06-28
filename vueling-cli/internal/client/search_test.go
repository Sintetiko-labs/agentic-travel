package client

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilterVuelingRows(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "testdata", "flightprice_mad_bcn.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var rows []vuelingFlightPriceRow
	if err := jsonUnmarshal(body, &rows); err != nil {
		t.Fatalf("decode: %v", err)
	}
	hits := filterVuelingRows(rows, "MAD", "BCN", "2026-07-01", "")
	if len(hits) == 0 {
		t.Fatal("expected flights for 2026-07-01")
	}
	if hits[0].FlightNumber == "" || hits[0].Price == "" {
		t.Fatalf("incomplete hit: %+v", hits[0])
	}
}
