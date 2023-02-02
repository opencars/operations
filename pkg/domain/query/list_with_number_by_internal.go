package query

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/opencars/seedwork"
	"github.com/opencars/translit"
)

type ListWithNumberByInternal struct {
	Number string
}

func (q *ListWithNumberByInternal) Prepare() {
	q.Number = translit.ToUA(strings.ToUpper(q.Number))
}

func (q *ListWithNumberByInternal) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.Number,
			validation.Required.Error(seedwork.Required),
			validation.Length(2, 18).Error(seedwork.Invalid),
		),
	)
}
