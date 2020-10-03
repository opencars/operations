package mapreduce

import "context"

type Mapper interface {
	Map(context.Context, <-chan []string, chan<- Entity) error
}
