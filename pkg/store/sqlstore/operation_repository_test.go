package sqlstore_test

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/opencars/operations/pkg/domain"
	"github.com/opencars/operations/pkg/store/sqlstore"
)

func TestOperationRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := domain.TestResource(t)
	assert.NoError(t, s.Resource().Create(context.Background(), resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	operation := domain.TestOperation(t)
	operation.ResourceID = resource.ID
	assert.NoError(t, s.Operation().Create(context.Background(), operation))
	assert.NotNil(t, operation)
}

func TestOperationRepository_DeleteByResourceID(t *testing.T) {
	resource := domain.TestResource(t)

	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	assert.NoError(t, s.Resource().Create(context.Background(), resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	operation := domain.TestOperation(t)
	operation.ResourceID = resource.ID
	assert.NoError(t, s.Operation().Create(context.Background(), operation))
	assert.NotNil(t, operation)

	affected, err := s.Operation().DeleteByResourceID(context.Background(), resource.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, affected)
}
