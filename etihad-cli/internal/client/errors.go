package client

import "fmt"

type APIError struct {
	Status int
	Body   string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("etihad api: HTTP %d: %s", e.Status, e.Body)
}
