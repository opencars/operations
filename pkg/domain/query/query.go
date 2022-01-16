package query

import (
	"github.com/opencars/schema"

	"github.com/opencars/operations/pkg/domain/model"
)

const (
	Ascending  string = "ASC"
	Descending string = "DESC"
)

var (
	source = schema.Source{
		Service: "operations",
		Version: "1.0",
	}
)

type Query interface {
	Prepare()
	Validate() error
}

func Process(q Query) error {
	q.Prepare()

	return model.Validate(q, "request")
}
