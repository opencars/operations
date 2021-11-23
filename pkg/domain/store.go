package domain

import (
	"context"

	"github.com/opencars/operations/pkg/domain/model"
)

type Store interface {
	Resource() ResourceRepository
	Operation() OperationRepository
}

type ResourceRepository interface {
	Create(ctx context.Context, resource *model.Resource) error
	Update(rctx context.Context, esource *model.Resource) error
	FindByUID(ctx context.Context, uid string) (*model.Resource, error)
	All(ctx context.Context) ([]model.Resource, error)
}

type OperationRepository interface {
	ReadOperationRepository
	WriteOperationRepository
}

type ReadOperationRepository interface {
	FindByNumber(ctx context.Context, number string, limit uint64, order string) ([]model.Operation, error)
	FindByVIN(ctx context.Context, vin string, limit uint64, order string) ([]model.Operation, error)
}

type WriteOperationRepository interface {
	Create(ctx context.Context, operations ...*model.Operation) error
	DeleteByResourceID(ctx context.Context, id int64) (int64, error)
}
