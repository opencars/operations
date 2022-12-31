package query

import (
	"github.com/opencars/schema"
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
