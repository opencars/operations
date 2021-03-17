package domain

import (
	"context"

	"github.com/opencars/operations/pkg/bulkreader"
)

type UserOperationService interface {
	FindByNumber(ctx context.Context, snumber string, limit uint64, order string) ([]Operation, error)
}

type Parser interface {
	Parse(ctx context.Context, resource *Resource, r *bulkreader.BulkReader) error
}
