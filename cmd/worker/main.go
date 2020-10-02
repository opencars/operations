package main

import (
	"flag"

	_ "github.com/lib/pq"
	"github.com/opencars/govdata"

	"github.com/opencars/operations/pkg/config"
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

	w := worker.New(store)

	modified, err := w.ModifiedResources()
	if err != nil {
		logger.Fatalf("modified: %v", err)
	}

	events := govdata.SubscribePackage(conf.Worker.PackageID, modified)
	for event := range events {
		if err := w.Process(event); err != nil {
			logger.Fatalf("process: %v", err)
		}
	}
}
