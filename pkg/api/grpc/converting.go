package grpc

import (
	"time"

	"github.com/opencars/grpc/pkg/common"
	"github.com/opencars/grpc/pkg/operation"

	"github.com/opencars/operations/pkg/domain"
)

func FromDomain(op *domain.Operation) *operation.Record {
	item := operation.Record{
		Number:  op.Number,
		Brand:   op.Brand,
		Model:   op.Model,
		Year:    int32(op.Year),
		Color:   op.Color,
		Kind:    op.Kind,
		Body:    op.Body,
		Purpose: op.Purpose,
		Department: &operation.Department{
			Code: op.DepCode,
			Name: op.Dep,
		},
	}

	if op.Capacity != nil {
		item.Capacity = int32(*op.Capacity)
	}

	if op.Fuel != nil {
		item.Fuel = *op.Fuel
	}

	if op.OwnWeight != nil {
		item.OwnWeight = int32(*op.OwnWeight)
	}

	if op.TotalWeight != nil {
		item.TotalWeight = int32(*op.TotalWeight)
	}

	if op.Date != "" {
		date, _ := time.Parse(domain.DateLayout, op.Date)
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
