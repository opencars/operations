package sqlstore

import (
	"database/sql"

	"github.com/opencars/operations/pkg/model"
	"github.com/opencars/operations/pkg/store"
)

// ResourceRepository is responsible for resources data.
type ResourceRepository struct {
	store *Store
}

func (r *ResourceRepository) Create(resource *model.Resource) error {
	rows, err := r.store.db.NamedQuery(
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

func (r *ResourceRepository) Update(resource *model.Resource) error {
	_, err := r.store.db.NamedQuery(
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

func (r *ResourceRepository) FindByUID(uid string) (*model.Resource, error) {
	var resource model.Resource

	err := r.store.db.Get(&resource,
		`SELECT id, uid, name, last_modified, url, created_at
		FROM resources
		WHERE uid = $1`,
		uid,
	)
	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func (r *ResourceRepository) All() ([]model.Resource, error) {
	resources := make([]model.Resource, 0)

	err := r.store.db.Select(&resources, `SELECT id, uid, name, last_modified, url, created_at FROM resources`)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
