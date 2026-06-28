package client

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilterBinterFlights(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "testdata", "search_lpa_tfn.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	hits, errMsg := parseBinterSearch(body, "LPA", "TFN", "2026-08-15")
	if errMsg != "" {
		t.Fatalf("parse error: %s", errMsg)
	}
	if len(hits) != 1 {
		t.Fatalf("expected 1 flight, got %d", len(hits))
	}
	if hits[0].FlightNumber != "NT101" || hits[0].Price != "63.00" {
		t.Fatalf("unexpected hit: %+v", hits[0])
	}
}
