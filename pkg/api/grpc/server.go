package grpc

import (
	"context"

	"github.com/opencars/grpc/pkg/operation"
)

type operationHandler struct {
	operation.UnimplementedServiceServer
	api *API
}

func (h *operationHandler) FindByNumber(ctx context.Context, r *operation.NumberRequest) (*operation.Response, error) {
	operations, err := h.api.svc.FindByNumber(ctx, r.Number, 0, "DESC")
	if err != nil {
		return nil, handleErr(err)
	}

	dto := operation.Response{
		Operations: make([]*operation.Record, 0, len(operations)),
	}

	for i := range operations {
		dto.Operations = append(dto.Operations, FromDomain(&operations[i]))
	}

	return &dto, nil
}
