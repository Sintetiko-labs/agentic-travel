// Package network enforces residential-only egress for travel CLIs.
package network

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var ProxyEnvVars = []string{
	"HTTP_PROXY", "HTTPS_PROXY", "ALL_PROXY",
	"http_proxy", "https_proxy", "all_proxy",
}

func SetProxyVars() []string {
	var set []string
	for _, k := range ProxyEnvVars {
		if strings.TrimSpace(os.Getenv(k)) != "" {
			set = append(set, k)
		}
	}
	return set
}

func EnsureResidential() {
	for _, k := range SetProxyVars() {
		fmt.Fprintf(os.Stderr, "warning: %s is set (%q) — travel CLIs ignore proxies; use Mac home IP, not VPN/datacenter egress\n",
			k, truncate(os.Getenv(k), 60))
	}
}

func ClearProxyEnv() {
	for _, k := range ProxyEnvVars {
		_ = os.Unsetenv(k)
	}
}

func PreprocessArgs(args []string) []string {
	if len(args) == 0 {
		EnsureResidential()
		return args
	}
	var out []string
	noProxy := false
	for _, a := range args {
		if a == "--no-proxy" {
			noProxy = true
			continue
		}
		out = append(out, a)
	}
	if noProxy {
		ClearProxyEnv()
		fmt.Fprintln(os.Stderr, "network: --no-proxy — cleared HTTP_PROXY/HTTPS_PROXY/ALL_PROXY")
	}
	EnsureResidential()
	return out
}

func DisableProxy(t *http.Transport) {
	if t == nil {
		return
	}
	t.Proxy = func(*http.Request) (*url.URL, error) { return nil, nil }
}

func DirectTransport() *http.Transport {
	t := http.DefaultTransport.(*http.Transport).Clone()
	DisableProxy(t)
	return t
}

func DirectClient(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: timeout, Transport: DirectTransport()}
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
