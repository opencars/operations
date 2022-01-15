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

func TestCustomerService_ListByNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []model.Operation{
		*model.TestOperation(t),
	}

	repo := mocks.NewMockOperationRepository(ctrl)
	repo.EXPECT().FindByNumber(gomock.Any(), expected[0].Number, uint64(10), query.Descending).Return(expected, nil)

	producer := mocks.NewMockProducer(ctrl)
	producer.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(nil)

	svc := service.NewCustomerService(repo, producer)

	q := query.ListByNumber{
		Number: expected[0].Number,
		UserID: "6585e1cf-d24e-4d9c-b012-453427498539",
	}

	actual, err := svc.ListByNumber(context.Background(), &q)
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, expected[0].Number, actual[0].Number)
	assert.Equal(t, expected[0].VIN, actual[0].VIN)
}

func TestCustomerService_ListByVIN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []model.Operation{
		*model.TestOperation(t),
	}

	repo := mocks.NewMockOperationRepository(ctrl)
	repo.EXPECT().FindByVIN(gomock.Any(), *expected[0].VIN, uint64(10), query.Descending).Return(expected, nil)

	producer := mocks.NewMockProducer(ctrl)
	producer.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(nil)

	svc := service.NewCustomerService(repo, producer)

	q := query.ListByVIN{
		VIN:    *expected[0].VIN,
		UserID: "6585e1cf-d24e-4d9c-b012-453427498539",
	}

	actual, err := svc.ListByVIN(context.Background(), &q)
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, expected[0].Number, actual[0].Number)
	assert.Equal(t, expected[0].VIN, actual[0].VIN)
}
