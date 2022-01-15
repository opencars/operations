package query

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/opencars/translit"

	"github.com/opencars/operations/pkg/domain/model"
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
			validation.Required.Error(model.Required),
			validation.Length(6, 18).Error(model.Invalid),
		),
	)
}
