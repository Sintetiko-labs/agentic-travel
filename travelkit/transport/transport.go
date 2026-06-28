package transport

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	utls "github.com/refraction-networking/utls"
)

// NewChromeTransport returns an http.Transport that mimics Chrome TLS (HTTP/1.1).
func NewChromeTransport() (*http.Transport, error) {
	if _, err := chromeSpecHTTP1(); err != nil {
		return nil, err
	}
	dial := func(ctx context.Context, network, addr string) (net.Conn, error) {
		if _, ok := ctx.Deadline(); !ok {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
			defer cancel()
		}
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		dialer := &net.Dialer{Timeout: 15 * time.Second}
		raw, err := dialer.DialContext(ctx, "tcp4", addr)
		if err != nil {
			raw, err = dialer.DialContext(ctx, network, addr)
			if err != nil {
				return nil, err
			}
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
		// CLIs must use the Mac's local network (residential IP). Do not route via
		// HTTP_PROXY / HTTPS_PROXY / ALL_PROXY — datacenter egress breaks Akamai WAF.
		Proxy: func(*http.Request) (*url.URL, error) { return nil, nil },
		DialTLSContext:        dial,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   15 * time.Second,
		ExpectContinueTimeout: time.Second,
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
