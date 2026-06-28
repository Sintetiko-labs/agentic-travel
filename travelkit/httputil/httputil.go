package httputil

import "io"

// Truncate shortens s to at most n bytes, adding an ellipsis when trimmed.
func Truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}

// ReadLimited reads up to n bytes from r.
func ReadLimited(r io.Reader, n int64) ([]byte, error) {
	return io.ReadAll(io.LimitReader(r, n))
}
