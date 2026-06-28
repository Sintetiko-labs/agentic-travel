package client

import (
	"io"
	"net/http"
	"strings"
)

func httpNewPost(url string, body []byte) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, url, strings.NewReader(string(body)))
}

func ioReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(io.LimitReader(r, 32<<20))
}
