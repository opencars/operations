package store

import (
	"errors"
)

var (
	// ErrRecordNotFound returned, when record was not found in the repository.
	ErrRecordNotFound = errors.New("record not found")
)

// Store is responsible for data manipulation.
type Store interface {
	Resource() ResourceRepository
	Operation() OperationRepository
}
