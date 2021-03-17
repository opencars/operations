package main

import (
	"context"
	"flag"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/opencars/operations/pkg/api/http"
	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/domain/user"
	"github.com/opencars/operations/pkg/logger"
	"github.com/opencars/operations/pkg/store/sqlstore"
)

func main() {
	cfg := flag.String("config", "config/config.yaml", "Path to the configuration file")
	port := flag.Int("port", 3000, "Port of the server")

	flag.Parse()

	conf, err := config.New(*cfg)
	if err != nil {
		logger.Fatalf("config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	svc := user.NewService(store.Operation())

	addr := ":" + strconv.Itoa(*port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Infof("Listening on %s...", addr)
	if err := http.Start(ctx, addr, &conf.Server, svc); err != nil {
		logger.Fatalf("http server failed: %v", err)
	}
}
