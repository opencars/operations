package parsing

import (
	"context"
	"fmt"

	"github.com/opencars/operations/pkg/domain"
	"github.com/opencars/operations/pkg/logger"
)

type reducer struct {
	repo domain.OperationRepository
}

func NewReducer(repo domain.OperationRepository) Reducer {
	return &reducer{
		repo: repo,
	}
}

func (r *reducer) Reduce(ctx context.Context, batches <-chan []Entity) error {
	for {
		select {
		case entities, ok := <-batches:
			if !ok {
				return nil
			}

			operations := make([]*domain.Operation, 0, len(batches))

			for i := range entities {
				operation, ok := entities[i].(*domain.Operation)
				if !ok {
					return fmt.Errorf("unexpected type: %T", entities[i])
				}

				operations = append(operations, operation)
			}

			if err := r.repo.Create(ctx, operations...); err != nil {
				return err
			}

			logger.Debugf("inserted %d entities", len(operations))
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
