package store

import (
	"github.com/opencars/operations/pkg/model"
)

type OperationRepository interface {
	Add(operation ...model.Operation) error
	Create(operation *model.Operation) error
	DeleteByResourceID(id int64) error
}

type ResourceRepository interface {
	Create(resource *model.Resource) error
	Update(resource *model.Resource) error
	FindByUID(uid string) (*model.Resource, error)
	All() ([]model.Resource, error)
}
