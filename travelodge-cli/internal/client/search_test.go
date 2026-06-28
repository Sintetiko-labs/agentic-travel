package client

import "testing"

func TestSplitDestinationTerms(t *testing.T) {
	terms := splitDestinationTerms("London & Manchester")
	if len(terms) != 2 || terms[0] != "London" || terms[1] != "Manchester" {
		t.Fatalf("got %#v", terms)
	}
	if got := splitDestinationTerms("London"); len(got) != 1 || got[0] != "London" {
		t.Fatalf("single: %#v", got)
	}
}

func TestSearchRejectsEmptyQuery(t *testing.T) {
	cl := New("")
	if _, err := cl.Search("   ", 1, 10); err == nil {
		t.Fatal("expected error for empty query")
	}
}
