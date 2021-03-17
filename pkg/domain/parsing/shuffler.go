package parsing

import "context"

type shuffler struct {
	size int
}

func NewShuffler(size int) Shuffler {
	return &shuffler{
		size: size,
	}
}

func (s *shuffler) Shuffle(ctx context.Context, entities <-chan Entity, batches chan<- []Entity) error {
	batch := make([]Entity, 0, s.size)

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
			batch = make([]Entity, 0, s.size)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
