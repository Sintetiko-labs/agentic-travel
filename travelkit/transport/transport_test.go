package transport

import (
	"testing"

	utls "github.com/refraction-networking/utls"
)

func TestChromeSpecHTTP1ALPN(t *testing.T) {
	spec, err := chromeSpecHTTP1()
	if err != nil {
		t.Fatal(err)
	}
	var alpn []string
	for _, ext := range spec.Extensions {
		if a, ok := ext.(*utls.ALPNExtension); ok {
			alpn = a.AlpnProtocols
		}
	}
	want := []string{"http/1.1"}
	if len(alpn) != len(want) || alpn[0] != want[0] {
		t.Fatalf("ALPN = %v, want %v", alpn, want)
	}
}

func TestNewChromeTransport(t *testing.T) {
	tr, err := NewChromeTransport()
	if err != nil {
		t.Fatal(err)
	}
	if tr == nil || tr.DialTLSContext == nil {
		t.Fatal("expected non-nil transport with DialTLSContext")
	}
	if tr.ForceAttemptHTTP2 {
		t.Fatal("HTTP/2 must be disabled")
	}
}
