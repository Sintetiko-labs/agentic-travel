package client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestNHSearchFixture(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "testdata", "search_madrid.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var resp nhSearchResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	resp.normalize()
	if len(resp.Data) != 2 {
		t.Fatalf("got %d hotels", len(resp.Data))
	}
	res := resp.toResult("Madrid", 1, 24, "NH Hotels", "https://www.nh-hotels.com")
	if res.Hotels[0].Name == "" || res.Hotels[0].HotelURL == "" {
		t.Fatalf("unexpected hit: %+v", res.Hotels[0])
	}
}
