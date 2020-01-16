package sqlstore_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/operations/pkg/model"
	"github.com/opencars/operations/pkg/store/sqlstore"
)

func TestResourceRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)
}

func TestResourceRepository_Update(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	resource.LastModified = time.Now().Add(-time.Minute)
	assert.NoError(t, s.Resource().Update(resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)
}

func TestResourceRepository_FindByUID(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	actual, err := s.Resource().FindByUID(resource.UID)
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestResourceRepository_All(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	actual, err := s.Resource().All()
	assert.NoError(t, err)
	assert.Len(t, actual, 1)
}
