package transport

import (
	"testing"
	utls "github.com/refraction-networking/utls"
)

func TestChromeSpecALPN(t *testing.T) {
	spec, err := chromeSpec()
	if err != nil { t.Fatal(err) }
	var alpn []string
	for _, ext := range spec.Extensions {
		if a, ok := ext.(*utls.ALPNExtension); ok { alpn = a.AlpnProtocols }
	}
	want := []string{"h2", "http/1.1"}
	if len(alpn) != len(want) || alpn[0] != want[0] || alpn[1] != want[1] {
		t.Fatalf("ALPN = %v, want %v", alpn, want)
	}
}

func TestSharedRoundTripperSingleton(t *testing.T) {
	a, err := SharedRoundTripper(); if err != nil { t.Fatal(err) }
	b, err := SharedRoundTripper(); if err != nil { t.Fatal(err) }
	if a != b { t.Fatal("expected singleton") }
}
