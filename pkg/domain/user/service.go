package user

import (
	"context"

	"github.com/opencars/translit"

	"github.com/opencars/operations/pkg/domain"
)

type Service struct {
	r domain.ReadOperationRepository
}

func NewService(r domain.ReadOperationRepository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) FindByNumber(ctx context.Context, number string, limit uint64, order string) ([]domain.Operation, error) {
	lexeme := translit.ToUA(number)

	operations, err := s.r.FindByNumber(ctx, lexeme, limit, order)
	if err != nil {
		return nil, err
	}

	for i := range operations {
		operations[i].Person = operations[i].PrettyPerson()
	}

	return operations, nil
}
