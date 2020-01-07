package teststore

import (
	"github.com/opencars/operations/pkg/model"
	"github.com/opencars/operations/pkg/store"
)

type ResourceRepository struct {
	store     *Store
	resources map[int64]*model.Resource
}

func (r *ResourceRepository) Create(resource *model.Resource) error {
	if resource.ID == 0 {
		resource.ID = int64(len(r.resources))
	}

	r.resources[resource.ID] = resource
	return nil
}

func (r *ResourceRepository) Update(resource *model.Resource) error {
	_, ok := r.resources[resource.ID]
	if !ok {
		return store.ErrRecordNotFound
	}

	r.resources[resource.ID] = resource
	return nil
}

func (r *ResourceRepository) FindByUID(uid string) (*model.Resource, error) {
	for _, v := range r.resources {
		if v.UID == uid {
			return v, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *ResourceRepository) All() ([]model.Resource, error) {
	resources := make([]model.Resource, 0)

	for _, v := range r.resources {
		resources = append(resources, *v)
	}

	return resources, nil
}
