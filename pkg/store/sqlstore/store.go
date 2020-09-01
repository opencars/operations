package sqlstore

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/store"
)

// Store is an implementation of store.Store interface based on SQL.
type Store struct {
	db                  *sqlx.DB
	operationRepository *OperationRepository
	resourceRepository  *ResourceRepository
}

// Resource returns repository, who is responsible for resources.
func (s *Store) Resource() store.ResourceRepository {
	if s.resourceRepository != nil {
		return s.resourceRepository
	}

	s.resourceRepository = &ResourceRepository{
		store: s,
	}

	return s.resourceRepository
}

// Operation returns repository, who is responsible for operations.
func (s *Store) Operation() store.OperationRepository {
	if s.operationRepository != nil {
		return s.operationRepository
	}

	s.operationRepository = &OperationRepository{
		store: s,
	}

	return s.operationRepository
}

// New returns new instance of Store.
func New(settings *config.Database) (*Store, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Name,
		settings.SSLMode,
		settings.Password,
	)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
