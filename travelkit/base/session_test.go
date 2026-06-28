package base

import (
	"testing"

	"github.com/fbelchi/travelkit/cookies"
	"github.com/fbelchi/travelkit/session"
)

func TestLoadPersistedCookies(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("MELIA_CONFIG_DIR", dir)

	if err := session.Save("MELIA", session.Data{Cookie: "sid=abc; _abck=1", BaseURL: "https://www.melia.com"}); err != nil {
		t.Fatal(err)
	}

	c := New("https://www.melia.com", "melia")
	if c.Cookie != "sid=abc; _abck=1" {
		t.Fatalf("cookie: got %q", c.Cookie)
	}
	if got := cookies.JarString(c.Jar, c.BaseURL); got == "" {
		t.Fatal("expected jar cookies")
	}
}

func TestEnvCookieOverridesFile(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("MELIA_CONFIG_DIR", dir)
	t.Setenv("MELIA_COOKIE", "from=env")

	_ = session.Save("MELIA", session.Data{Cookie: "from=file"})

	c := New("https://www.melia.com", "melia")
	if c.Cookie != "from=env" {
		t.Fatalf("got %q want from=env", c.Cookie)
	}
}

func TestSavePersistedCookies(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("NH_CONFIG_DIR", dir)

	c := New("https://www.nh-hotels.com", "nh")
	c.ApplyCookieHeader("token=xyz")
	if err := c.SavePersistedCookies(); err != nil {
		t.Fatal(err)
	}
	d, err := session.Load("NH")
	if err != nil {
		t.Fatal(err)
	}
	if d.Cookie != "token=xyz" {
		t.Fatalf("got %q", d.Cookie)
	}
}
