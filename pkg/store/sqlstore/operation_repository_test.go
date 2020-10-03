package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/operations/pkg/model"
	"github.com/opencars/operations/pkg/store/sqlstore"
)

func TestOperationRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	resource := model.TestResource(t)
	assert.NoError(t, s.Resource().Create(resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	operation := model.TestOperation(t)
	operation.ResourceID = resource.ID
	assert.NoError(t, s.Operation().Create(operation))
	assert.NotNil(t, operation)
}

func TestOperationRepository_DeleteByResourceID(t *testing.T) {
	resource := model.TestResource(t)

	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("operations", "resources")

	assert.NoError(t, s.Resource().Create(resource))
	assert.NotNil(t, resource)
	assert.EqualValues(t, 1, resource.ID)

	operation := model.TestOperation(t)
	operation.ResourceID = resource.ID
	assert.NoError(t, s.Operation().Create(operation))
	assert.NotNil(t, operation)

	affected, err := s.Operation().DeleteByResourceID(resource.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, affected)
}
