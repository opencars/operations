package handler

import "net/http"

var (
	// ErrInvalidOrder returned, when order parameter is not ASC or DESC.
	ErrInvalidOrder = NewError(http.StatusBadRequest, "params.order.is_not_valid")
	// ErrInvalidLimit returned, when limit parameter can not be parsed into uint64.
	ErrInvalidLimit = NewError(http.StatusBadRequest, "params.limit.is_not_valid")
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents error with http status code.
type StatusError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Message
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// NewError creates new error instance.
func NewError(code int, message string) Error {
	return &StatusError{
		Code:    code,
		Message: message,
	}
}
