package mapreduce

import "context"

type Reducer interface {
	Reduce(context.Context, <-chan []Entity) error
}
