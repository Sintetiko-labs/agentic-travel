package client

import "fmt"

type APIError struct {
	Status int
	Body   string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("asiana api: HTTP %d: %s", e.Status, e.Body)
}
