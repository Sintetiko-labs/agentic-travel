package client

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fbelchi/travelkit/akamai"
	tkbase "github.com/fbelchi/travelkit/base"
	"github.com/fbelchi/travelkit/transport"
)

const BaseURL = "https://www.marriott.com"

// Client talks to Marriott public endpoints.
type Client struct {
	*tkbase.Client
	Brand string
}

// Brands supported by this CLI (shared parent API).
var Brands = []string{
	"Marriott",
	"Marriott Hotels",
	"JW Marriott",
	"The Ritz-Carlton",
	"St. Regis",
	"W Hotels",
	"Edition",
	"Luxury Collection",
	"Westin",
	"Sheraton",
	"Le Méridien",
	"Renaissance Hotels",
	"Autograph Collection",
	"Tribute Portfolio",
	"AC Hotels",
	"AC Hotels by Marriott",
	"Aloft",
	"Moxy",
	"Courtyard by Marriott",
	"Residence Inn",
}

func New(brand string) *Client {
	c := &Client{Client: tkbase.New(BaseURL, "marriott"), Brand: brand}
	c.enableChromeTLS()
	return c
}

type chromeRoundTripper struct {
	port     int
	fallback http.RoundTripper
}

func (t *chromeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return transport.FetchViaChrome(t.port, req, t.fallback)
}

// enableChromeTLS routes HTTP through headed Chrome CDP when Akamai session cookies
// are present and remote debugging is listening.
func (c *Client) enableChromeTLS() {
	if c.Cookie == "" || !akamai.SessionReady(c.Cookie) {
		return
	}
	if os.Getenv("MARRIOTT_STD_HTTP") == "1" {
		return
	}
	port := marriottChromePort()
	if !transport.CDPAvailable(port) {
		return
	}
	var fallback http.RoundTripper
	if c.HTTP != nil {
		fallback = c.HTTP.Transport
	}
	c.HTTP.Transport = &chromeRoundTripper{port: port, fallback: fallback}
}

func marriottChromePort() int {
	if p := strings.TrimSpace(os.Getenv("MARRIOTT_CHROME_PORT")); p != "" {
		var n int
		if _, err := fmt.Sscanf(p, "%d", &n); err == nil && n > 0 {
			return n
		}
	}
	return transport.CDPPortFromEnv()
}
