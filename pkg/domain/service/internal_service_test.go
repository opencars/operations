package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/opencars/operations/pkg/domain/mocks"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/domain/query"
	"github.com/opencars/operations/pkg/domain/service"
)

func TestInternalService_ListByNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []model.Operation{
		*model.TestOperation(t),
	}

	repo := mocks.NewMockOperationRepository(ctrl)
	repo.EXPECT().FindByNumber(gomock.Any(), expected[0].Number, uint64(100), query.Descending).Return(expected, nil)

	svc := service.NewInternalService(repo)
	actual, err := svc.ListByNumber(context.Background(), &query.ListWithNumberByInternal{Number: expected[0].Number})
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, expected[0].Number, actual[0].Number)
	assert.Equal(t, expected[0].VIN, actual[0].VIN)
}

func TestInternalService_ListByVIN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []model.Operation{
		*model.TestOperation(t),
	}

	repo := mocks.NewMockOperationRepository(ctrl)
	repo.EXPECT().FindByVIN(gomock.Any(), *expected[0].VIN, uint64(100), query.Descending).Return(expected, nil)

	svc := service.NewInternalService(repo)
	actual, err := svc.ListByVIN(context.Background(), &query.ListWithVINByInternal{VIN: *expected[0].VIN})
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, expected[0].Number, actual[0].Number)
	assert.Equal(t, expected[0].VIN, actual[0].VIN)
}
