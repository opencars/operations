package sqlstore

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/domain"
)

// Store is an implementation of domain.Store interface based on SQL.
type Store struct {
	db *sqlx.DB

	operationOnce       sync.Once
	operationRepository *OperationRepository

	resourceOnce       sync.Once
	resourceRepository *ResourceRepository
}

// Resource returns repository, who is responsible for resources.
func (s *Store) Resource() domain.ResourceRepository {
	s.operationOnce.Do(func() {
		s.resourceRepository = &ResourceRepository{
			store: s,
		}
	})

	return s.resourceRepository
}

// Operation returns repository, who is responsible for operations.
func (s *Store) Operation() domain.OperationRepository {
	s.resourceOnce.Do(func() {
		s.operationRepository = &OperationRepository{
			store: s,
		}
	})

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
