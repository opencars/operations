package grpc

import (
	"context"

	"github.com/opencars/grpc/pkg/operation"

	"github.com/opencars/operations/pkg/domain/query"
)

type operationHandler struct {
	operation.UnimplementedServiceServer
	api *API
}

func (h *operationHandler) FindByNumber(ctx context.Context, r *operation.NumberRequest) (*operation.Response, error) {
	q := query.ListWithNumberByInternal{
		Number: r.Number,
	}

	operations, err := h.api.svc.ListByNumber(ctx, &q)
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

func (h *operationHandler) FindByVIN(ctx context.Context, r *operation.VINRequest) (*operation.Response, error) {
	q := query.ListWithVINByInternal{
		VIN: r.Vin,
	}

	operations, err := h.api.svc.ListByVIN(ctx, &q)
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
