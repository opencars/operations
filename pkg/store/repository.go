package store

import (
	"github.com/opencars/operations/pkg/model"
)

// OperationRepository is responsible for operations data.
type OperationRepository interface {
	Create(operation ...model.Operation) error
	DeleteByResourceID(int64) (int64, error)
}

// ResourceRepository is responsible for resources data.
type ResourceRepository interface {
	Create(resource *model.Resource) error
	Update(resource *model.Resource) error
	FindByUID(uid string) (*model.Resource, error)
	All() ([]model.Resource, error)
}
