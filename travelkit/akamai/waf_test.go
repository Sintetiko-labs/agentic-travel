package akamai

import "testing"

func TestIsIncapsulaChallenge(t *testing.T) {
	body := `<html><script src="/_Incapsula_Resource?SWJIYLWA=abc"></script></html>`
	if !IsIncapsulaChallenge(body) {
		t.Fatal("expected incapsula challenge")
	}
}

func TestIsWAFBlockedIncapsula200(t *testing.T) {
	body := `<html><script src="/_Incapsula_Resource?SWJIYLWA=abc"></script></html>`
	if !IsWAFBlocked(200, body) {
		t.Fatal("expected 200 incapsula to be blocked")
	}
}

func TestProbeOKJSON(t *testing.T) {
	if !ProbeOK(200, `{"ok":true}`) {
		t.Fatal("expected JSON 200 to be OK")
	}
	if ProbeOK(200, `<html>challenge</html>`) {
		t.Fatal("expected HTML 200 to fail probe")
	}
}
