package worker

import (
	"context"

	"github.com/opencars/operations/pkg/mapreduce"
)

type Shuffler struct {
	size int
}

func NewShuffler(size int) mapreduce.Shuffler {
	return &Shuffler{
		size: size,
	}
}

func (s *Shuffler) Shuffle(ctx context.Context, entities <-chan mapreduce.Entity, batches chan<- []mapreduce.Entity) error {
	batch := make([]mapreduce.Entity, 0, s.size)

	for {
		select {
		case entity, ok := <-entities:
			if !ok {
				if len(batch) != 0 {
					batches <- batch
				}
				return nil
			}

			if len(batch) < s.size {
				batch = append(batch, entity)
				continue
			}

			batches <- batch

			// TODO: Find another way to find too much memory allocation.
			batch = make([]mapreduce.Entity, 0, s.size)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
