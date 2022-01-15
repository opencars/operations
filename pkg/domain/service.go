package domain

import (
	"context"
	"io"

	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/domain/query"
)

type CustomerService interface {
	FindByNumber(context.Context, *query.ListByNumber) ([]model.Operation, error)
	FindByVIN(context.Context, *query.ListByVIN) ([]model.Operation, error)
}

type InternalService interface {
	FindByNumber(context.Context, *query.ListWithNumberByInternal) ([]model.Operation, error)
	FindByVIN(context.Context, *query.ListWithVINByInternal) ([]model.Operation, error)
}

type Parser interface {
	Parse(ctx context.Context, resource *model.Resource, rc io.ReadCloser) error
}
