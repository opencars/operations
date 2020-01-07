package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/store"
)

type Store struct {
	db                  *sqlx.DB
	operationRepository *OperationRepository
	resourceRepository  *ResourceRepository
}

func (s *Store) Resource() store.ResourceRepository {
	if s.resourceRepository != nil {
		return s.resourceRepository
	}

	s.resourceRepository = &ResourceRepository{
		store: s,
	}

	return s.resourceRepository
}

func (s *Store) Operation() store.OperationRepository {
	if s.operationRepository != nil {
		return s.operationRepository
	}

	s.operationRepository = &OperationRepository{
		store: s,
	}

	return s.operationRepository
}

func New(conf *config.Settings) (*Store, error) {
	var info string
	if conf.DB.Password == "" {
		info = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Name,
		)
	} else {
		info = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name,
		)
	}

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &Store{
		db: db,
	}, nil
}
