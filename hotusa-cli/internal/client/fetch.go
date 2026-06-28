package client

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	tkbase "github.com/fbelchi/travelkit/base"
	utls "github.com/refraction-networking/utls"
)

const hotusaTLSServerName = "www.hotusa.com"

func (c *Client) fetchHotusaHTML(path string) (string, error) {
	if path == "" {
		path = "/"
	}
	dialer := &net.Dialer{Timeout: 15 * time.Second}
	tr := &http.Transport{
		DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			raw, err := dialer.DialContext(ctx, network, "www.hotusa.com:443")
			if err != nil {
				return nil, err
			}
			cfg := &utls.Config{ServerName: hotusaTLSServerName, InsecureSkipVerify: true}
			uconn := utls.UClient(raw, cfg, utls.HelloChrome_131)
			if err := uconn.HandshakeContext(ctx); err != nil {
				raw.Close()
				return nil, err
			}
			return uconn, nil
		},
	}
	client := &http.Client{Transport: tr, Timeout: c.HTTP.Timeout}
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return "", err
	}
	req.Host = "www.hotusa.com"
	c.SetDocumentHeaders(req)
	c.ApplyCookie(req)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	text := string(body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", &tkbase.HTTPError{Status: resp.StatusCode, Body: tkbase.Truncate(text, 300)}
	}
	return text, nil
}

func hotusaPaths(query string) []string {
	slug := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(query), " ", "-"))
	paths := []string{"/es/", "/es/hoteles/", "/"}
	if slug != "" {
		paths = append([]string{"/es/hoteles/" + slug + "/", "/es/destinos/" + slug + "/"}, paths...)
	}
	return paths
}

func hotusaSearchErr(err error) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	if strings.Contains(msg, "certificate") || strings.Contains(msg, "tls") {
		return fmt.Errorf("hotusa TLS mismatch (cert is %s) — run: hotusa session chrome --wait", hotusaTLSServerName)
	}
	if he, ok := err.(*tkbase.HTTPError); ok && he.Status == 400 {
		return fmt.Errorf("hotusa blocked request — run: hotusa session chrome --wait (HTTP %d)", he.Status)
	}
	return err
}
