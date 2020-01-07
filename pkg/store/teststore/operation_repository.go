package teststore

import (
	"github.com/opencars/operations/pkg/model"
)

type OperationRepository struct {
	store      *Store
	operations []model.Operation
}

func (r *OperationRepository) Add(operations ...model.Operation) error {
	r.operations = append(r.operations, operations...)
	return nil
}

func (r *OperationRepository) Create(operation *model.Operation) error {
	r.operations = append(r.operations, *operation)
	return nil
}

func (r *OperationRepository) DeleteByResourceID(id int64) error {
	operations := make([]model.Operation, 0)
	for _, op := range r.operations {
		if op.ResourceID != id {
			operations = append(operations, op)
		}
	}
	r.operations = operations

	return nil
}
