package client

import (
	"strings"
	"testing"
)

func TestMarriottSearchURL(t *testing.T) {
	path := marriottSearchURL("London")
	if !strings.Contains(path, "destinationAddress.city=London") {
		t.Fatalf("city: %s", path)
	}
	if !strings.Contains(path, "destinationAddress.country=GB") {
		t.Fatalf("country: %s", path)
	}
}

func TestDefaultSearchProbePath(t *testing.T) {
	if !strings.HasPrefix(DefaultSearchProbePath(), "/search/findHotels.mi") {
		t.Fatalf("got %q", DefaultSearchProbePath())
	}
}

func TestSearchRejectsEmptyQuery(t *testing.T) {
	cl := New("")
	if _, err := cl.Search("   ", 1, 10); err == nil {
		t.Fatal("expected error for empty query")
	}
}

func TestMarriottBlockedErr(t *testing.T) {
	noCookie := marriottBlockedErr("").Error()
	if !strings.Contains(noCookie, "session chrome") {
		t.Fatalf("no cookie: %q", noCookie)
	}
	withCookie := marriottBlockedErr("_abck=x; bm_sz=y").Error()
	if !strings.Contains(withCookie, "stale") {
		t.Fatalf("with cookie: %q", withCookie)
	}
}
