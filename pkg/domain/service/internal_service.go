package service

import (
	"context"

	"github.com/opencars/operations/pkg/domain"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/domain/query"
	"github.com/opencars/operations/pkg/logger"
)

type InternalService struct {
	r  domain.ReadOperationRepository
	kd domain.KoatuuDecoder
}

func NewInternalService(r domain.ReadOperationRepository, kd domain.KoatuuDecoder) *InternalService {
	return &InternalService{
		r:  r,
		kd: kd,
	}
}

func (s *InternalService) ListByNumber(ctx context.Context, q *query.ListWithNumberByInternal) ([]model.Operation, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	operations, err := s.r.FindByNumber(ctx, q.Number, 100, query.Descending)
	if err != nil {
		return nil, err
	}

	if s.kd == nil {
		return operations, nil
	}

	exist := make(map[string][]int, 0)
	codes := make([]string, 0)
	for i := range operations {
		code := operations[i].RegAddress

		if code == nil {
			continue
		}

		if _, ok := exist[*code]; ok {
			continue
		}

		exist[*code] = append(exist[*code], i)
		codes = append(codes, *code)
	}

	if len(codes) > 0 {
		results, err := s.kd.Decode(ctx, codes...)
		if err != nil {
			logger.Errorf("failed to decode koatuu codes: %+v", codes)
		} else {
			for i, code := range codes {
				for _, v := range exist[code] {
					if len(results) <= i {
						logger.Errorf("unexpected length of koatuu response: %+v", code)
						continue
					}

					if results[i].Error != nil {
						logger.Errorf("failed to decode koatuu code: %+v", code)
						continue
					}

					operations[v].FullRegAddress = &results[i].Summary
				}
			}
		}
	}

	return operations, nil
}

func (s *InternalService) ListByVIN(ctx context.Context, q *query.ListWithVINByInternal) ([]model.Operation, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	operations, err := s.r.FindByVIN(ctx, q.VIN, 100, query.Descending)
	if err != nil {
		return nil, err
	}

	return operations, nil
}
