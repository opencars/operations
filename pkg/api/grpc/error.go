package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/opencars/operations/pkg/domain"
)

var (
	ErrNotFound = status.Error(codes.NotFound, "record.not_found")
)

func handleErr(err error) error {
	switch err {
	case domain.ErrNotFound:
		return ErrNotFound
	default:
		return err
	}
}
