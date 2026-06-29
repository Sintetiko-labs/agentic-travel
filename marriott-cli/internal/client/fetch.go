package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	tkbase "github.com/fbelchi/travelkit/base"
)

// DefaultSearchProbePath is the findHotels URL used by session doctor.
func DefaultSearchProbePath() string {
	return marriottSearchURL("London")
}

// fetchSearchHTML GETs a Marriott search page with browser-like navigation headers.
func (c *Client) fetchSearchHTML(path string) (string, error) {
	c.Throttle()
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return "", err
	}
	c.SetDocumentHeaders(req)
	req.Header.Set("referer", c.BaseURL+"/search/default.mi")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	c.ApplyCookie(req)
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	text := string(body)
	if resp.StatusCode == 403 && c.ChromeFetchEnabled() {
		chromeBody, status, ferr := c.FetchViaChromeReq(req)
		if ferr == nil {
			text = string(chromeBody)
			if status >= 200 && status < 300 {
				return text, nil
			}
			return "", &tkbase.HTTPError{Status: status, Body: tkbase.Truncate(text, 300)}
		}
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", &tkbase.HTTPError{Status: resp.StatusCode, Body: tkbase.Truncate(text, 300)}
	}
	return text, nil
}

func marriottSearchURLWithDates(city string, from, to time.Time) string {
	return fmt.Sprintf("/search/findHotels.mi?destinationAddress.city=%s&destinationAddress.country=GB&roomCount=1&numAdultsPerRoom=2&lengthOfStay=1&fromDate=%s&toDate=%s&deviceType=desktop-web&view=list",
		strings.ReplaceAll(strings.TrimSpace(city), " ", "+"),
		from.Format("01/02/2006"),
		to.Format("01/02/2006"))
}
