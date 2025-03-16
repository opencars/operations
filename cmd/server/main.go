package main

import (
	"context"
	"flag"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/opencars/schema/client"
	"github.com/opencars/seedwork/logger"
	"golang.org/x/sync/errgroup"

	"github.com/opencars/operations/pkg/api/grpc"
	"github.com/opencars/operations/pkg/api/http"
	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/domain/service"
	"github.com/opencars/operations/pkg/koatuu"
	"github.com/opencars/operations/pkg/store/sqlstore"
)

func main() {
	cfg := flag.String("config", "config/config.yaml", "Path to the configuration file")
	httpPort := flag.Int("http-port", 8080, "Port for HTTP server")
	grpcPort := flag.Int("grpc-port", 3000, "Port for gRPC server")

	flag.Parse()

	conf, err := config.New(*cfg)
	if err != nil {
		logger.Fatalf("config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	c, err := client.New(conf.NATS.Address())
	if err != nil {
		logger.Fatalf("nats: %v", err)
	}

	producer, err := c.Producer()
	if err != nil {
		logger.Fatalf("producer: %v", err)
	}

	// Initialize store
	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	kd, err := koatuu.NewService(conf.GRPC.KOATUU.Address())
	if err != nil {
		logger.Errorf("koatuu service: %s", err)
	}

	// Create services
	customerSvc := service.NewCustomerService(store.Operation(), producer)
	internalSvc := service.NewInternalService(store.Operation(), kd)

	// Create context with cancellation
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create errgroup with cancellation
	g, ctx := errgroup.WithContext(ctx)

	// Start HTTP server
	g.Go(func() error {
		addr := ":" + strconv.Itoa(*httpPort)
		logger.Infof("Starting HTTP server on %s...", addr)
		return http.Start(ctx, addr, &conf.Server, customerSvc)
	})

	// Start gRPC server
	g.Go(func() error {
		addr := ":" + strconv.Itoa(*grpcPort)
		logger.Infof("Starting gRPC server on %s...", addr)
		api := grpc.New(addr, internalSvc)
		return api.Run(ctx)
	})

	// Wait for interrupt signal or error from servers
	logger.Infof("Servers started successfully. Press Ctrl+C to stop...")
	if err := g.Wait(); err != nil {
		logger.Fatalf("Server error: %v", err)
	}
	logger.Infof("Servers stopped gracefully")
}
