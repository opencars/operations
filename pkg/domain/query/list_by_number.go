package query

import (
	"strconv"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/opencars/schema"
	"github.com/opencars/schema/vehicle"
	"github.com/opencars/translit"

	"github.com/opencars/operations/pkg/domain/model"
)

type ListByNumber struct {
	UserID  string
	TokenID string
	Number  string
	Limit   string
	Order   string
}

func (q *ListByNumber) Prepare() {
	q.Order = strings.ToUpper(q.Order)
	q.Number = translit.ToUA(strings.ToUpper(q.Number))
}

func (q *ListByNumber) GetLimit() uint64 {
	if q.Limit == "" {
		return 10
	}

	num, err := strconv.ParseInt(q.Limit, 10, 64)
	if err != nil {
		panic(err)
	}

	if num < 0 {
		return 10
	}

	return uint64(num)
}

func (q *ListByNumber) GetOrder() string {
	if q.Order == "" {
		return Descending
	}

	return q.Order
}

func (q *ListByNumber) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.UserID,
			validation.Required.Error(model.Required),
		),
		validation.Field(
			&q.Number,
			validation.Required.Error(model.Required),
			validation.Length(6, 18).Error(model.Invalid),
		),
		validation.Field(
			&q.Limit,
			is.Int.Error(model.IsNotInreger),
		),
		validation.Field(
			&q.Order,
			validation.In(Ascending, Descending).Error(model.Invalid),
		),
	)
}

func (q *ListByNumber) Event(operations ...model.Operation) schema.Producable {
	msg := vehicle.OperationSearched{
		UserId:       q.UserID,
		TokenId:      q.TokenID,
		Number:       q.Number,
		ResultAmount: uint32(len(operations)),
		SearchedAt:   timestamppb.New(time.Now().UTC()),
	}

	return schema.New(&source, &msg).Message(
		schema.WithSubject(schema.OperationCustomerActions),
	)
}
