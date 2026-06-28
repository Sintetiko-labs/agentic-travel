package client

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDecodeMeliaSearchFixture(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "testdata", "search_madrid.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	resp, err := decodeMeliaSearch(body)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(resp.Hotels) != 2 || resp.Hotels[0].Name != "Meliá Castilla" {
		t.Fatalf("unexpected hotels: %+v", resp.Hotels)
	}
	res := resp.toResult("Madrid", 1, 24, "Meliá", "https://www.melia.com")
	if res.Total != 2 || res.Hotels[0].HotelURL == "" {
		t.Fatalf("unexpected result: %+v", res)
	}
}

func TestShouldFallbackSearch(t *testing.T) {
	if !shouldFallbackSearch(errString("search: HTTP 404: not found")) {
		t.Fatal("expected 404 fallback")
	}
	if shouldFallbackSearch(errString("akamai blocked")) {
		t.Fatal("should not fallback on akamai block")
	}
}

type errString string

func (e errString) Error() string { return string(e) }
