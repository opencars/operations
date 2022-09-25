package service

import (
	"context"

	"github.com/opencars/operations/pkg/domain"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/domain/query"
)

type InternalService struct {
	r domain.ReadOperationRepository
}

func NewInternalService(r domain.ReadOperationRepository) *InternalService {
	return &InternalService{
		r: r,
	}
}

func (s *InternalService) ListByNumber(ctx context.Context, q *query.ListWithNumberByInternal) ([]model.Operation, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	operations, err := s.r.FindByNumber(ctx, q.Number, 100, query.Descending)
	if err != nil {
		return nil, err
	}

	return operations, nil
}

func (s *InternalService) ListByVIN(ctx context.Context, q *query.ListWithVINByInternal) ([]model.Operation, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	operations, err := s.r.FindByVIN(ctx, q.VIN, 100, query.Descending)
	if err != nil {
		return nil, err
	}

	return operations, nil
}
