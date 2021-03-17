package domain

import "context"

type Store interface {
	Resource() ResourceRepository
	Operation() OperationRepository
}

type ResourceRepository interface {
	Create(ctx context.Context, resource *Resource) error
	Update(rctx context.Context, esource *Resource) error
	FindByUID(ctx context.Context, uid string) (*Resource, error)
	All(ctx context.Context) ([]Resource, error)
}

type OperationRepository interface {
	ReadOperationRepository
	WriteOperationRepository
}

type ReadOperationRepository interface {
	FindByNumber(ctx context.Context, number string, limit uint64, order string) ([]Operation, error)
}

type WriteOperationRepository interface {
	Create(ctx context.Context, operations ...*Operation) error
	DeleteByResourceID(ctx context.Context, id int64) (int64, error)
}
