package transport

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
)

var errProtocolNegotiated = errors.New("protocol negotiated")

var ChromeHelloID = utls.HelloChrome_131

var (
	sharedOnce sync.Once
	sharedRT   http.RoundTripper
	sharedErr  error
)

func SharedRoundTripper() (http.RoundTripper, error) {
	sharedOnce.Do(func() {
		sharedRT, sharedErr = newChromeRoundTripper()
	})
	return sharedRT, sharedErr
}

func NewChromeTransport() (http.RoundTripper, error) {
	return SharedRoundTripper()
}

type chromeRoundTripper struct {
	sync.Mutex
	cachedConnections map[string]net.Conn
	cachedTransports  map[string]http.RoundTripper
	dialer            *net.Dialer
}

func newChromeRoundTripper() (http.RoundTripper, error) {
	if _, err := chromeSpec(); err != nil {
		return nil, err
	}
	return &chromeRoundTripper{
		cachedConnections: make(map[string]net.Conn),
		cachedTransports:  make(map[string]http.RoundTripper),
		dialer:            &net.Dialer{Timeout: 15 * time.Second},
	}, nil
}

func (rt *chromeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	addr := dialTLSAddr(req)
	if _, ok := rt.cachedTransports[addr]; !ok {
		if err := rt.ensureTransport(req, addr); err != nil {
			return nil, err
		}
	}
	return rt.cachedTransports[addr].RoundTrip(req)
}

func (rt *chromeRoundTripper) ensureTransport(req *http.Request, addr string) error {
	switch strings.ToLower(req.URL.Scheme) {
	case "http":
		rt.Lock()
		defer rt.Unlock()
		rt.cachedTransports[addr] = &http.Transport{
			Proxy: noProxy, DialContext: rt.dialer.DialContext,
			MaxIdleConns: 100, MaxIdleConnsPerHost: 10, IdleConnTimeout: 90 * time.Second,
		}
		return nil
	case "https":
	default:
		return fmt.Errorf("invalid URL scheme: %q", req.URL.Scheme)
	}
	_, err := rt.dialTLS(req.Context(), "tcp", addr)
	switch err {
	case errProtocolNegotiated:
		return nil
	case nil:
		panic("dialTLS returned no error when determining cachedTransports")
	default:
		return err
	}
}

func (rt *chromeRoundTripper) dialTLS(ctx context.Context, networkName, addr string) (net.Conn, error) {
	rt.Lock()
	defer rt.Unlock()
	if conn := rt.cachedConnections[addr]; conn != nil {
		delete(rt.cachedConnections, addr)
		return conn, nil
	}
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
		defer cancel()
	}
	raw, err := rt.dialer.DialContext(ctx, "tcp4", addr)
	if err != nil {
		raw, err = rt.dialer.DialContext(ctx, networkName, addr)
		if err != nil {
			return nil, err
		}
	}
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		host = addr
	}
	spec, err := chromeSpec()
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
	if rt.cachedTransports[addr] != nil {
		return u, nil
	}
	switch u.ConnectionState().NegotiatedProtocol {
	case http2.NextProtoTLS:
		rt.cachedTransports[addr] = &http2.Transport{
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				return rt.dialTLS(context.Background(), network, addr)
			},
		}
	default:
		rt.cachedTransports[addr] = &http.Transport{
			Proxy: noProxy, DialTLSContext: rt.dialTLS,
			MaxIdleConns: 100, MaxIdleConnsPerHost: 10, IdleConnTimeout: 90 * time.Second,
			TLSHandshakeTimeout: 15 * time.Second, ExpectContinueTimeout: time.Second,
		}
	}
	rt.cachedConnections[addr] = u
	return nil, errProtocolNegotiated
}

func dialTLSAddr(req *http.Request) string {
	host, port, err := net.SplitHostPort(req.URL.Host)
	if err == nil {
		return net.JoinHostPort(host, port)
	}
	return net.JoinHostPort(req.URL.Host, "443")
}

func noProxy(*http.Request) (*url.URL, error) { return nil, nil }

func chromeSpec() (utls.ClientHelloSpec, error) { return utls.UTLSIdToSpec(ChromeHelloID) }
