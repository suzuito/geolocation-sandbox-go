package web

import "github.com/suzuito/geolocation-sandbox-go/store"

import "fmt"

// HTTPError ...
type HTTPError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// FromStoreError ...
func FromStoreError(err store.Error) *HTTPError {
	return &HTTPError{
		Message: err.Raw().Error(),
		Status:  500,
	}
}

// NewNewStoreCoreError ...
func NewNewStoreCoreError(err error) *HTTPError {
	return &HTTPError{
		Status:  500,
		Message: fmt.Sprintf("%s", err.Error()),
	}
}
