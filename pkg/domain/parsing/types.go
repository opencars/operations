package parsing

import (
	"context"

	"github.com/opencars/operations/pkg/csv"

	"github.com/opencars/operations/pkg/domain/model"
)

type Convertible interface {
	Convert() *model.Operation
}

type Mapper interface {
	Map(context.Context, *model.Resource, *csv.Reader, chan<- []Convertible) error
}

type Reducer interface {
	Reduce(context.Context, <-chan []model.Operation) error
}

type Shuffler interface {
	Shuffle(context.Context, <-chan []Convertible, chan<- []model.Operation) error
}
