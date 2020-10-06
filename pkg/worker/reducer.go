package worker

import (
	"context"
	"fmt"

	"github.com/opencars/operations/pkg/logger"
	"github.com/opencars/operations/pkg/mapreduce"
	"github.com/opencars/operations/pkg/model"
	"github.com/opencars/operations/pkg/store"
)

type Reducer struct {
	store store.Store
}

func NewReducer(s store.Store) mapreduce.Reducer {
	return &Reducer{
		store: s,
	}
}

func (r *Reducer) Reduce(ctx context.Context, batches <-chan []mapreduce.Entity) error {
	for {
		select {
		case entities, ok := <-batches:
			if !ok {
				return nil
			}

			operations := make([]*model.Operation, 0, len(batches))

			for i := range entities {
				operation, ok := entities[i].(*model.Operation)
				if !ok {
					return fmt.Errorf("unexpected type: %T", entities[i])
				}

				operations = append(operations, operation)
			}

			if err := r.store.Operation().Create(operations...); err != nil {
				return err
			}

			logger.Debugf("inserted %d entities", len(operations))
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
