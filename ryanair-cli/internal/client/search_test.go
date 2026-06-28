package client

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFarfndFixture(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "testdata", "farfnd_stn_dub.json"))
	if err != nil {
		t.Fatal(err)
	}
	res, err := parseFarfndResponse(b, "https://www.ryanair.com", "STN", "DUB", "2026-07-01", 1, 5)
	if err != nil {
		t.Fatal(err)
	}
	_ = b
	if res.Total < 1 {
		t.Fatalf("expected fares, got total=%d", res.Total)
	}
	if res.Source != "farfnd" {
		t.Fatalf("source=%q", res.Source)
	}
	if res.Flights[0].Price == "" {
		t.Fatalf("missing price: %+v", res.Flights[0])
	}
}
