package portal

import (
	"context"
	"errors"
	"io"

	"github.com/opencars/operations/pkg/csv"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/domain/parsing"
)

type Mapper struct {
	size int
}

func NewMapper(size int) *Mapper {
	return &Mapper{
		size: size,
	}
}

func (m *Mapper) Map(ctx context.Context, resource *model.Resource, r *csv.Reader, convertibles chan<- []parsing.Convertible) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			records := make([]Operation, 0)

			err := r.ReadBulk(m.size, &records)
			if err == nil || errors.Is(err, io.EOF) {
				cc := make([]parsing.Convertible, 0)
				for i := range records {
					cc = append(cc, &records[i])
				}

				convertibles <- cc
			}

			if errors.Is(err, io.EOF) {
				return nil
			}

			if err != nil {
				return err
			}
		}
	}
}
