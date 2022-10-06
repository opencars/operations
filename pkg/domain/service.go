package domain

import (
	"context"
	"io"

	"github.com/opencars/grpc/pkg/koatuu"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/domain/query"
)

type CustomerService interface {
	ListByNumber(context.Context, *query.ListByNumber) ([]model.Operation, error)
	ListByVIN(context.Context, *query.ListByVIN) ([]model.Operation, error)
}

type InternalService interface {
	ListByNumber(context.Context, *query.ListWithNumberByInternal) ([]model.Operation, error)
	ListByVIN(context.Context, *query.ListWithVINByInternal) ([]model.Operation, error)
}

type Parser interface {
	Parse(ctx context.Context, resource *model.Resource, rc io.ReadCloser) error
}

type KoatuuDecoder interface {
	Decode(context.Context, ...string) ([]*koatuu.DecodeResultItem, error)
}
