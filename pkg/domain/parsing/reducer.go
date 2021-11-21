package parsing

import (
	"context"
	"fmt"

	"github.com/opencars/operations/pkg/domain"
	"github.com/opencars/operations/pkg/domain/model"
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

func (r *reducer) Reduce(ctx context.Context, batches <-chan []model.Operation) error {
	for {
		select {
		case entities, ok := <-batches:
			if !ok {
				return nil
			}

			operations := make([]*model.Operation, 0, len(batches))

			for i := range entities {
				operations = append(operations, &entities[i])
			}

			if err := r.repo.Create(ctx, operations...); err != nil {
				fmt.Println(err)
				return err
			}

			logger.Debugf("inserted %d entities", len(operations))
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
