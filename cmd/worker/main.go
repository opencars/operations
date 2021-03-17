package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/domain/parsing"
	"github.com/opencars/operations/pkg/logger"
	"github.com/opencars/operations/pkg/store/sqlstore"
	"github.com/opencars/operations/pkg/worker"
)

func main() {
	var path string

	flag.StringVar(&path, "config", "./config/config.yaml", "Path to the configuration file")

	flag.Parse()

	conf, err := config.New(path)
	if err != nil {
		logger.Fatalf("config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	parser := parsing.NewMapReduce().
		WithMapper(parsing.NewMapper()).
		WithReducer(parsing.NewReducer(store.Operation())).
		WithsShuffler(parsing.NewShuffler(conf.Worker.InsertBatchSize))

	w := worker.New(store, parser)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := w.Process(ctx, conf.Worker.PackageID); err != nil {
		logger.Fatalf("process: %v", err)
	}
}
