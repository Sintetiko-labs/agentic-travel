package session

import (
	"testing"
	"time"
)

func TestSaveLoadRoundTrip(t *testing.T) {
	t.Setenv("MELIA_CONFIG_DIR", t.TempDir())
	want := Data{Cookie: "foo=bar; _abck=xyz", BaseURL: "https://www.melia.com"}
	if err := Save("MELIA", want); err != nil {
		t.Fatal(err)
	}
	got, err := Load("MELIA")
	if err != nil {
		t.Fatal(err)
	}
	if got.Cookie != want.Cookie {
		t.Fatalf("cookie: got %q want %q", got.Cookie, want.Cookie)
	}
	if got.BaseURL != want.BaseURL {
		t.Fatalf("base_url: got %q want %q", got.BaseURL, want.BaseURL)
	}
	if got.Captured.IsZero() {
		t.Fatal("expected captured_at to be set")
	}
	if time.Since(got.Captured) > time.Minute {
		t.Fatal("captured_at too old")
	}
}

func TestFilePath(t *testing.T) {
	t.Setenv("MELIA_CONFIG_DIR", "/tmp/melia-test")
	if p := FilePath("MELIA"); p != "/tmp/melia-test/cookies.json" {
		t.Fatalf("got %q", p)
	}
}
