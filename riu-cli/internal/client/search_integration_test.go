package client

import (
	"strings"
	"testing"

	"github.com/fbelchi/travelkit/parse"
)

func TestParseRIUMadridFixture(t *testing.T) {
	c := New("")
	html, err := c.FetchHTML(c.BaseURL + "/es/hotels/europa/espana/madrid")
	if err != nil {
		t.Skip("network:", err)
	}
	if !strings.Contains(html, "ng-state") {
		t.Fatalf("missing ng-state, len=%d", len(html))
	}
	if !strings.Contains(html, "KEY_HOTELS_FILTER") {
		t.Fatalf("ng-state present but missing KEY_HOTELS_FILTER, len=%d", len(html))
	}
	rows := parse.ParseRIUNgState(html, c.BaseURL, "")
	if len(rows) == 0 {
		t.Fatalf("expected hotels from Madrid page (html len=%d)", len(html))
	}
}
