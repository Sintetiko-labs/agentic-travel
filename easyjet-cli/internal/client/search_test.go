package client

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEasyjetResponseToResult(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "testdata", "ejavailability_mad_pmi.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var resp easyjetResponse
	if err := jsonUnmarshal(body, &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	res := resp.toResult("MAD", "PMI", "2026-07-05", "", 1, 24, "easyJet")
	if res.Total == 0 || len(res.Flights) == 0 {
		t.Fatalf("expected flights, got %+v", res)
	}
	if res.Flights[0].Price == "" {
		t.Fatalf("expected price on first flight: %+v", res.Flights[0])
	}
}
