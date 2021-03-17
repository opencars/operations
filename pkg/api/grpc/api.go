package grpc

import (
	"context"
	"net"

	"github.com/opencars/grpc/pkg/operation"
	"google.golang.org/grpc"

	"github.com/opencars/operations/pkg/domain"
)

// API represents gRPC API server.
type API struct {
	Addr string
	s    *grpc.Server
	svc  domain.UserOperationService
}

func New(addr string, svc domain.UserOperationService) *API {
	return &API{
		Addr: addr,
		svc:  svc,
	}
}

func (a *API) Run(ctx context.Context) error {
	errs := make(chan error)
	go func() {
		errs <- a.run()
	}()

	select {
	case <-ctx.Done():
		return a.close()
	case err := <-errs:
		return err
	}
}

func (a *API) run() error {
	listener, err := net.Listen("tcp", a.Addr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			RequestLoggingInterceptor,
			// ErrorInterceptor,
		),
	}

	a.s = grpc.NewServer(opts...)
	operation.RegisterServiceServer(a.s, &operationHandler{api: a})

	return a.s.Serve(listener)
}

// Close gracefully stops grpc API server.
func (a *API) close() error {
	if a.s != nil {
		a.s.GracefulStop()
	}

	return nil
}
