package parsing

import (
	"context"

	"github.com/opencars/operations/pkg/domain"
)

type Entity interface{}

type Mapper interface {
	Map(context.Context, *domain.Resource, <-chan []string, chan<- Entity) error
}

type Reducer interface {
	Reduce(context.Context, <-chan []Entity) error
}

type Shuffler interface {
	Shuffle(context.Context, <-chan Entity, chan<- []Entity) error
}
