package service

import (
	"context"

	"github.com/opencars/schema"

	"github.com/opencars/operations/pkg/domain"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/domain/query"
)

type CustomerService struct {
	r domain.ReadOperationRepository
	p schema.Producer
}

func NewCustomerService(r domain.ReadOperationRepository, p schema.Producer) *CustomerService {
	return &CustomerService{
		r: r,
		p: p,
	}
}

func (s *CustomerService) ListByNumber(ctx context.Context, q *query.ListByNumber) ([]model.Operation, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	operations, err := s.r.FindByNumber(ctx, q.Number, q.GetLimit(), q.GetOrder())
	if err != nil {
		return nil, err
	}

	for i := range operations {
		operations[i].Person = operations[i].PrettyPerson()
	}

	if err := s.p.Produce(ctx, q.Event(operations...)); err != nil {
		return nil, err
	}

	return operations, nil
}

func (s *CustomerService) ListByVIN(ctx context.Context, q *query.ListByVIN) ([]model.Operation, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	operations, err := s.r.FindByVIN(ctx, q.VIN, q.GetLimit(), q.GetOrder())
	if err != nil {
		return nil, err
	}

	for i := range operations {
		operations[i].Person = operations[i].PrettyPerson()
	}

	if err := s.p.Produce(ctx, q.Event(operations...)); err != nil {
		return nil, err
	}

	return operations, nil
}
