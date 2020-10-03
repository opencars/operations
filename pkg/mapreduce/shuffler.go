package mapreduce

import "context"

type Shuffler interface {
	Shuffle(context.Context, <-chan Entity, chan<- []Entity) error
}
