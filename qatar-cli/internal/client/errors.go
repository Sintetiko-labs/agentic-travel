package client

import "fmt"

type APIError struct {
	Status int
	Body   string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("qatar api: HTTP %d: %s", e.Status, e.Body)
}
