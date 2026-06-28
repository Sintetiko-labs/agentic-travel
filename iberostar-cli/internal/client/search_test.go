package client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestIberostarSearchFixture(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("..", "..", "testdata", "search_madrid.json"))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var resp struct {
		Data struct {
			SearchHotels struct {
				Total  int `json:"total"`
				Hotels []struct {
					ID, Name, City string
				} `json:"hotels"`
			} `json:"searchHotels"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if resp.Data.SearchHotels.Total != 2 || resp.Data.SearchHotels.Hotels[0].Name == "" {
		t.Fatalf("unexpected payload: %+v", resp.Data.SearchHotels)
	}
}
