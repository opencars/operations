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

	flag.StringVar(&path, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	// Get configuration.
	conf, err := config.New(path)
	if err != nil {
		logger.Fatal(err)
	}

	// Register store.
	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatal(err)
	}

	w := worker.New(store)
	events := govdata.SubscribePackage(conf.Worker.PackageID, w.ModifiedResources())
	for event := range events {
		if err := w.Process(event); err != nil {
			logger.Fatal(err)
		}
	}
}
