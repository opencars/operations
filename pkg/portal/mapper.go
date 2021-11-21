package portal

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/opencars/operations/pkg/csv"
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

func (m *Mapper) Map(ctx context.Context, r *csv.Reader, convertibles chan<- []parsing.Convertible) error {
	var decoder *csv.RowDecoder

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			records := make([]parsing.Convertible, 0, m.size)

			for i := 0; i < m.size; i++ {
				row, err := r.Read()
				if errors.Is(err, io.EOF) {
					convertibles <- records
					return nil
				}

				if err != nil {
					return err
				}

				if decoder == nil {
					fields := make(map[string]int)

					for j, f := range row {
						fields[strings.ToUpper(f)] = j
					}

					decoder = csv.NewRowDecoder(fields)

					i--
					continue
				}

				var obj Operation
				if err := decoder.Decode(row, &obj); err != nil {
					log.Print(row, err)
				}

				records = append(records, &obj)
			}

			convertibles <- records
		}
	}
}
