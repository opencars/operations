package main

import (
	"context"
	"flag"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/opencars/operations/pkg/api/grpc"
	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/domain/service"
	"github.com/opencars/operations/pkg/koatuu"
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

	kd, err := koatuu.NewService(conf.GRPC.KOATUU.Address())
	if err != nil {
		logger.Errorf("koatuu service: %s", err)
	}

	svc := service.NewInternalService(store.Operation(), kd)

	addr := ":" + strconv.Itoa(*port)
	api := grpc.New(addr, svc)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger.Infof("Listening on %s...", addr)
	if err := api.Run(ctx); err != nil {
		logger.Fatalf("grpc: %v", err)
	}
}
