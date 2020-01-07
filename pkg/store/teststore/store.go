package teststore

import (
	"github.com/opencars/operations/pkg/model"
	"github.com/opencars/operations/pkg/store"
)

type Store struct {
	operationRepository *OperationRepository
	resourceRepository  *ResourceRepository
}

func (s *Store) Vehicle() store.OperationRepository {
	if s.operationRepository != nil {
		return s.operationRepository
	}

	s.operationRepository = &OperationRepository{
		store:      s,
		operations: make([]model.Operation, 0),
	}

	return s.operationRepository
}

func (s *Store) Resource() store.ResourceRepository {
	if s.resourceRepository != nil {
		return s.resourceRepository
	}

	s.resourceRepository = &ResourceRepository{
		store:     s,
		resources: make(map[int64]*model.Resource),
	}

	return s.resourceRepository
}

func New() *Store {
	return &Store{}
}
