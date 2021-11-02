package parsing

import (
	"context"

	"github.com/opencars/operations/pkg/domain/model"
)

type shuffler struct{}

func NewShuffler() Shuffler {
	return &shuffler{}
}

func (s *shuffler) Shuffle(ctx context.Context, convertibles <-chan []Convertible, batches chan<- []model.Operation) error {
	batch := make([]model.Operation, 0)

	for {
		select {
		case cc, ok := <-convertibles:
			if !ok {
				return nil
			}

			for _, c := range cc {
				batch = append(batch, *c.Convert())
			}

			batches <- batch

			// TODO: Find another way to find too much memory allocation.
			batch = make([]model.Operation, 0)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
