package transport

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	utls "github.com/refraction-networking/utls"
)

const defaultDialTimeout = 10 * time.Second
const defaultTLSHandshake = 15 * time.Second
const defaultResponseHeader = 20 * time.Second

// NewChromeTransport returns an http.Transport that mimics Chrome TLS over HTTP/1.1.
// HTTP/2 is intentionally disabled: the prior h2+uTLS round-tripper could hang
// without honoring http.Client timeouts (see docs/SMOKE_MAC_HOTELS_ES BLOCKED).
func NewChromeTransport() (*http.Transport, error) {
	if _, err := chromeSpecHTTP1(); err != nil {
		return nil, err
	}
	dial := func(ctx context.Context, network, addr string) (net.Conn, error) {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		raw, err := (&net.Dialer{Timeout: defaultDialTimeout}).DialContext(ctx, network, addr)
		if err != nil {
			return nil, err
		}
		spec, err := chromeSpecHTTP1()
		if err != nil {
			raw.Close()
			return nil, err
		}
		u := utls.UClient(raw, &utls.Config{ServerName: host}, utls.HelloCustom)
		if err := u.ApplyPreset(&spec); err != nil {
			raw.Close()
			return nil, fmt.Errorf("utls preset: %w", err)
		}
		if err := u.HandshakeContext(ctx); err != nil {
			raw.Close()
			return nil, fmt.Errorf("utls handshake: %w", err)
		}
		return u, nil
	}
	return &http.Transport{
		DialTLSContext:        dial,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   defaultTLSHandshake,
		ResponseHeaderTimeout: defaultResponseHeader,
		ExpectContinueTimeout: time.Second,
		ForceAttemptHTTP2:     false,
	}, nil
}

func chromeSpecHTTP1() (utls.ClientHelloSpec, error) {
	spec, err := utls.UTLSIdToSpec(utls.HelloChrome_Auto)
	if err != nil {
		return spec, err
	}
	for _, ext := range spec.Extensions {
		if a, ok := ext.(*utls.ALPNExtension); ok {
			a.AlpnProtocols = []string{"http/1.1"}
		}
	}
	return spec, nil
}
