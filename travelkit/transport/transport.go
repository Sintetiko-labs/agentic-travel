package transport

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	utls "github.com/refraction-networking/utls"

	"github.com/fbelchi/travelkit/network"
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
	tr := &http.Transport{
		DialTLSContext:        dial,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   15 * time.Second,
		ExpectContinueTimeout: time.Second,
	}
	network.DisableProxy(tr)
	return tr, nil
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
