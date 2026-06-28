package client

import "testing"

func TestHiltonLocationPath(t *testing.T) {
	if got := hiltonLocationPath(""); got != "" {
		t.Fatalf("empty: %q", got)
	}
	if got := hiltonLocationPath("St. James's"); got != "/en/locations/united-kingdom/st-jamess/" {
		t.Fatalf("special chars: %q", got)
	}
	if got := hiltonLocationPath("London"); got != "/en/locations/united-kingdom/london/" {
		t.Fatalf("london: %q", got)
	}
}

func TestSearchRejectsEmptyQuery(t *testing.T) {
	cl := New("")
	if _, err := cl.Search("   ", 1, 10); err == nil {
		t.Fatal("expected error for empty query")
	}
}
