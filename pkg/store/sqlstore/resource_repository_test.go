package sqlstore_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/store/sqlstore"
)

func TestResourceRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(context.Background(), resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)
}

func TestResourceRepository_Update(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(context.Background(), resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	resource.LastModified = time.Now().Add(-time.Minute)
	assert.NoError(t, s.Resource().Update(context.Background(), resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)
}

func TestResourceRepository_FindByUID(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(context.Background(), resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	actual, err := s.Resource().FindByUID(context.Background(), resource.UID)
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestResourceRepository_All(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(context.Background(), resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	actual, err := s.Resource().All(context.Background())
	assert.NoError(t, err)
	assert.Len(t, actual, 1)
}
