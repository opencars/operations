package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/opencars/operations/pkg/api/http"
	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/logger"
	"github.com/opencars/operations/pkg/store/sqlstore"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.yaml", "Path to the configuration file")

	flag.Parse()

	conf, err := config.New(configPath)
	if err != nil {
		logger.Fatalf("config: %v", err)
	}

	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addr := ":8080"
	logger.Infof("Listening on %s...", addr)
	if err := http.Start(ctx, addr, &conf.Server, store.Operation()); err != nil {
		logger.Fatalf("http server failed: %v", err)
	}
}
