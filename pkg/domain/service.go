package domain

import (
	"context"
	"io"

	"github.com/opencars/operations/pkg/domain/model"
)

type UserOperationService interface {
	FindByNumber(ctx context.Context, number string, limit uint64, order string) ([]model.Operation, error)
	FindByVIN(ctx context.Context, vin string, limit uint64, order string) ([]model.Operation, error)
}

type Parser interface {
	Parse(ctx context.Context, resource *model.Resource, rc io.ReadCloser) error
}
