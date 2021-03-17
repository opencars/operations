package parsing

import (
	"context"

	"github.com/opencars/operations/pkg/domain"
)

type mapper struct{}

func NewMapper() Mapper { return &mapper{} }

func (m *mapper) Map(ctx context.Context, resource *domain.Resource, rows <-chan []string, entities chan<- Entity) error {
	for {
		select {
		case msg, ok := <-rows:
			if !ok {
				return nil
			}

			op, err := domain.OperationFromGov(msg)
			if err != nil {
				return err
			}
			op.ResourceID = resource.ID

			entities <- op
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
