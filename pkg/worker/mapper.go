package worker

import (
	"context"

	"github.com/opencars/operations/pkg/mapreduce"
	"github.com/opencars/operations/pkg/model"
)

type Mapper struct {
	resource *model.Resource
}

func NewMapper(resource *model.Resource) mapreduce.Mapper {
	return &Mapper{
		resource: resource,
	}
}

func (m *Mapper) Map(ctx context.Context, rows <-chan []string, entities chan<- mapreduce.Entity) error {
	for {
		select {
		case msg, ok := <-rows:
			if !ok {
				return nil
			}

			op, err := model.OperationFromGov(msg)
			if err != nil {
				return err
			}
			op.ResourceID = m.resource.ID

			entities <- op
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
