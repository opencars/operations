package koatuu

import (
	"context"

	"google.golang.org/grpc"

	"github.com/opencars/grpc/pkg/koatuu"
)

type Service struct {
	c koatuu.ServiceClient
}

func NewService(addr string) (*Service, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Service{
		c: koatuu.NewServiceClient(conn),
	}, nil
}

func (s *Service) Decode(ctx context.Context, codes ...string) ([]*koatuu.DecodeResultItem, error) {
	req := koatuu.DecodeRequest{
		Codes: codes,
	}

	resp, err := s.c.Decode(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp.Items, nil
}
