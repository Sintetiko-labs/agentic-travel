package akamai

import "testing"

func TestSessionReadyAkamai(t *testing.T) {
	if !SessionReady("_abck=1; bm_sz=2") {
		t.Fatal("expected akamai session ready")
	}
}

func TestSessionReadyAnalyticsOnly(t *testing.T) {
	c := "dtCookie=1; rxVisitor=2; didomi_token=3"
	if SessionReady(c) {
		t.Fatal("analytics-only cookies should not be session-ready")
	}
	if !HasSessionMaterial(c) {
		t.Fatal("expected session material")
	}
}

func TestNeedsAkamaiWAF(t *testing.T) {
	if !NeedsAkamaiWAF("nh") {
		t.Fatal("nh should need akamai")
	}
	if NeedsAkamaiWAF("melia") {
		t.Fatal("melia should not require akamai pair")
	}
}

func TestSiteSessionReady(t *testing.T) {
	if !SiteSessionReady("melia", "OptanonConsent=1; dtCookie=abc") {
		t.Fatal("melia site cookies should be ready")
	}
	if SiteSessionReady("melia", "foo=bar") {
		t.Fatal("generic cookie should not satisfy melia site check")
	}
	if !SiteSessionReady("iberostar", "_session_id=abc; foo=bar") {
		t.Fatal("iberostar session id should be ready")
	}
}
