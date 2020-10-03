package mapreduce

import (
	"context"
	"io"

	"golang.org/x/sync/errgroup"

	"github.com/opencars/operations/pkg/bulkreader"
	"github.com/opencars/operations/pkg/logger"
)

const (
	DefaultMappers   = 100
	DefaultReducers  = 2
	DefaultShufflers = 2
	DefaultBulkSize  = 1000
)

type MapReduce struct {
	shuffler Shuffler
	mapper   Mapper
	reducer  Reducer

	shufflers int
	mappers   int
	reducers  int
	bulkSize  int

	rows     chan []string
	entities chan Entity
	batches  chan []Entity

	r *bulkreader.BulkReader
}

func NewMapReduce(r *bulkreader.BulkReader) *MapReduce {
	return &MapReduce{
		rows:     make(chan []string),
		entities: make(chan Entity),
		batches:  make(chan []Entity),

		reducers:  DefaultReducers,
		mappers:   DefaultMappers,
		shufflers: DefaultShufflers,
		bulkSize:  DefaultBulkSize,

		r: r,
	}
}

func (mr *MapReduce) Process(ctx context.Context) (resErr error) {
	reducerGroup, reducerCtx := errgroup.WithContext(context.Background())
	for i := 0; i < mr.reducers; i++ {
		logger.Debugf("starting %d reducer", i)
		reducerGroup.Go(func() error {
			return mr.reducer.Reduce(reducerCtx, mr.batches)
		})
	}
	defer func() {
		logger.Debugf("waiting for reducers")
		if err := reducerGroup.Wait(); err != nil {
			resErr = err
		}
	}()

	shufflerGroup, shufflerCtx := errgroup.WithContext(context.Background())
	for i := 0; i < mr.shufflers; i++ {
		logger.Debugf("starting %d shuffler", i)
		shufflerGroup.Go(func() error {
			return mr.shuffler.Shuffle(shufflerCtx, mr.entities, mr.batches)
		})
	}
	defer func() {
		logger.Debugf("waiting for shufflers")
		if err := shufflerGroup.Wait(); err != nil {
			resErr = err
		}

		logger.Debugf("closing batches channel")
		close(mr.batches)
	}()

	mapperGroup, mapperCtx := errgroup.WithContext(context.Background())
	for i := 0; i < mr.mappers; i++ {
		logger.Debugf("starting %d mapper", i)
		mapperGroup.Go(func() error {
			return mr.mapper.Map(mapperCtx, mr.rows, mr.entities)
		})
	}
	defer func() {
		logger.Debugf("waiting for mappers")
		if err := mapperGroup.Wait(); err != nil {
			resErr = err
		}

		logger.Debugf("closing entities channel")
		close(mr.entities)
	}()

	logger.Debugf("starting mapperDispatcher")

	defer func() {
		logger.Debugf("closing rows channel")
		close(mr.rows)
	}()
	if err := mr.mapperDispatcher(ctx); err != nil {
		return err
	}

	return nil
}

func (mr *MapReduce) mapperDispatcher(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Debugf("mapperDispatcher stopped")
			return ctx.Err()
		default:
			logger.Debugf("reading rows")
			messages, err := mr.r.ReadBulk(mr.bulkSize)

			if err == nil || err == io.EOF {
				for _, msg := range messages {
					mr.rows <- msg
				}
			}

			if err == io.EOF {
				return nil
			}

			if err != nil {
				return err
			}
		}
	}
}
