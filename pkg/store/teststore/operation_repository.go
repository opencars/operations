package teststore

import (
	"github.com/opencars/operations/pkg/model"
)

// OperationRepository is responsible for operations data.
type OperationRepository struct {
	store      *Store
	operations []model.Operation
}

// Create adds new records to the operations table.
func (r *OperationRepository) Create(operations ...model.Operation) error {
	r.operations = append(r.operations, operations...)
	return nil
}

// FindByNumber returns list operations on verhicles with specified number plates.
func (r *OperationRepository) FindByNumber(number string, limit uint64, order string) ([]model.Operation, error) {
	operations := make([]model.Operation, 0)

	for _, op := range r.operations {
		if op.Number == number {
			operations = append(operations, op)
		}
	}

	return operations, nil
}

// DeleteByResourceID removes records with specified resource_id from operations table.
func (r *OperationRepository) DeleteByResourceID(id int64) (int64, error) {
	was := len(r.operations)

	operations := make([]model.Operation, 0)
	for _, op := range r.operations {
		if op.ResourceID != id {
			operations = append(operations, op)
		}
	}
	r.operations = operations

	return int64(was - len(r.operations)), nil
}
