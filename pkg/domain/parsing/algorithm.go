package parsing

import (
	"context"
	"io"

	"golang.org/x/sync/errgroup"

	"github.com/opencars/operations/pkg/csv"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/logger"
)

const (
	DefaultMappers   = 1
	DefaultReducers  = 10
	DefaultShufflers = 10
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

	convertibles chan []Convertible
	batches      chan []model.Operation
}

func NewMapReduce() *MapReduce {
	return &MapReduce{
		reducers:  DefaultReducers,
		shufflers: DefaultShufflers,
		mappers:   DefaultMappers,
		bulkSize:  DefaultBulkSize,
	}
}

func (mr *MapReduce) Parse(ctx context.Context, resource *model.Resource, rc io.ReadCloser) (resErr error) {
	mr.convertibles = make(chan []Convertible)
	mr.batches = make(chan []model.Operation)

	csvReader := csv.NewReader(rc, ';')

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
			return mr.shuffler.Shuffle(shufflerCtx, resource, mr.convertibles, mr.batches)
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

	mapperGroup, mapperCtx := errgroup.WithContext(ctx)
	for i := 0; i < mr.mappers; i++ {
		logger.Debugf("starting %d mapper", i)
		mapperGroup.Go(func() error {
			return mr.mapper.Map(mapperCtx, csvReader, mr.convertibles)
		})
	}
	defer func() {
		logger.Debugf("waiting for mappers")
		if err := mapperGroup.Wait(); err != nil {
			resErr = err
		}

		logger.Debugf("closing convertibles")
		close(mr.convertibles)
	}()

	return nil
}
