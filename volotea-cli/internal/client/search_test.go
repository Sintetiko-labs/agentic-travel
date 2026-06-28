package client

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilterVoloteaJourneys(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "testdata", "search_bcn_ovd.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	hits := parseVoloteaSearch(body, "BCN", "OVD", "2026-07-23", "")
	if len(hits) != 1 {
		t.Fatalf("expected 1 flight, got %d", len(hits))
	}
	if hits[0].FlightNumber != "V73573" || hits[0].Price != "104.86" {
		t.Fatalf("unexpected hit: %+v", hits[0])
	}
}
