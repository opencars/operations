package query

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/opencars/seedwork"
	"github.com/opencars/translit"
)

type ListWithVINByInternal struct {
	VIN string
}

func (q *ListWithVINByInternal) Prepare() {
	q.VIN = translit.ToLatin(strings.ToUpper(q.VIN))
}

func (q *ListWithVINByInternal) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.VIN,
			validation.Required.Error(seedwork.Required),
			validation.Length(6, 18).Error(seedwork.Invalid),
		),
	)
}
