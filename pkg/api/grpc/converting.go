package grpc

import (
	"time"

	"github.com/opencars/grpc/pkg/common"
	"github.com/opencars/grpc/pkg/operation"

	"github.com/opencars/operations/pkg/domain/model"
)

func FromDomain(op *model.Operation) *operation.Record {
	item := operation.Record{
		Number:      op.Number,
		Brand:       op.Brand,
		Model:       op.Model,
		Year:        int32(op.Year),
		Capacity:    int32(op.Capacity),
		Color:       op.Color,
		Kind:        op.Kind,
		Body:        op.Body,
		Purpose:     op.Purpose,
		OwnWeight:   int32(op.OwnWeight),
		TotalWeight: int32(op.TotalWeight),
		Department: &operation.Department{
			Code: op.DepCode,
			Name: op.Dep,
		},
	}

	if op.Fuel != nil {
		item.Fuel = *op.Fuel
	}

	if op.Date != "" {
		date, _ := time.Parse(model.DateLayout, op.Date)
		item.Date = &common.Date{
			Year:  int32(date.Year()),
			Month: int32(date.Month()),
			Day:   int32(date.Day()),
		}
	}

	item.Owner = &operation.Owner{
		Entity: operation.Owner_UNKNOWN,
		Registration: &operation.Owner_Territory{
			Code: int32(op.RegCode),
		},
	}

	switch op.Person {
	case "J":
		item.Owner.Entity = operation.Owner_LEGAL
	case "P":
		item.Owner.Entity = operation.Owner_INDIVIDUAL
	}

	return &item
}
