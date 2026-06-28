package network

import (
	"net/http"
	"os"
	"testing"
)

func TestSetProxyVars(t *testing.T) {
	t.Setenv("HTTP_PROXY", "http://proxy:8080")
	t.Setenv("HTTPS_PROXY", "")
	got := SetProxyVars()
	if len(got) != 1 || got[0] != "HTTP_PROXY" {
		t.Fatalf("SetProxyVars: got %v", got)
	}
}

func TestClearProxyEnv(t *testing.T) {
	t.Setenv("HTTP_PROXY", "http://proxy:8080")
	ClearProxyEnv()
	for _, k := range ProxyEnvVars {
		if os.Getenv(k) != "" {
			t.Fatalf("%s still set", k)
		}
	}
}

func TestPreprocessArgsNoProxy(t *testing.T) {
	t.Setenv("HTTPS_PROXY", "http://proxy:3128")
	args := PreprocessArgs([]string{"travelodge", "--no-proxy", "search", "London"})
	if len(args) != 3 || args[1] != "search" {
		t.Fatalf("args: %v", args)
	}
}

func TestDirectTransportIgnoresProxy(t *testing.T) {
	tr := DirectTransport()
	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
	u, err := tr.Proxy(req)
	if u != nil || err != nil {
		t.Fatalf("Proxy: url=%v err=%v", u, err)
	}
}
