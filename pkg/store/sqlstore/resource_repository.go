package sqlstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/opencars/operations/pkg/domain"
)

// ResourceRepository is responsible for resources data.
type ResourceRepository struct {
	store *Store
}

func (r *ResourceRepository) Create(ctx context.Context, resource *domain.Resource) error {
	rows, err := r.store.db.NamedQueryContext(ctx,
		`INSERT INTO resources
		(
			uid, name, last_modified, url
		)
		VALUES
		(
			:uid, :name, :last_modified, :url
		)
		ON CONFLICT(uid) DO UPDATE SET last_modified = :last_modified RETURNING id`,
		resource,
	)

	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&resource.ID); err != nil {
			return err
		}
	}

	return nil
}

func (r *ResourceRepository) Update(ctx context.Context, resource *domain.Resource) error {
	_, err := r.store.db.NamedQueryContext(ctx,
		`UPDATE resources SET
			uid = :uid, name = :name, last_modified = :last_modified, url = :url
		WHERE id = :id`,
		resource,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ResourceRepository) FindByUID(ctx context.Context, uid string) (*domain.Resource, error) {
	var resource domain.Resource

	err := r.store.db.GetContext(ctx, &resource,
		`SELECT id, uid, name, last_modified, url, created_at
		FROM resources
		WHERE uid = $1`,
		uid,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func (r *ResourceRepository) All(ctx context.Context) ([]domain.Resource, error) {
	resources := make([]domain.Resource, 0)

	err := r.store.db.SelectContext(ctx, &resources,
		`SELECT id, uid, name, last_modified, url, created_at
		FROM resources`,
	)

	if err != nil {
		return nil, err
	}

	return resources, nil
}
