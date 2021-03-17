package http

import (
	"net/http"

	"github.com/opencars/httputil"
)

var (
	// ErrInvalidOrder returned, when order parameter is not ASC or DESC.
	ErrInvalidOrder = httputil.NewError(http.StatusBadRequest, "params.order.is_not_valid")
	// ErrInvalidLimit returned, when limit parameter can not be parsed into uint64.
	ErrInvalidLimit = httputil.NewError(http.StatusBadRequest, "params.limit.is_not_valid")
)
