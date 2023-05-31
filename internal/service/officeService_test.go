package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
	mock_repository "github.com/tumbleweedd/mediasoft-intership/customer-service/internal/repository/mocks"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"testing"
)

func TestOfficeService_CreateOffice(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockOffice, expectedUUID uuid.UUID, expectedOffice model.Office)

	testTable := []struct {
		name           string
		req            *customer.CreateOfficeRequest
		mockBehavior   mockBehavior
		expectedUUID   uuid.UUID
		expectedOffice model.Office
		expectedError  error
		expectedResult *customer.CreateOfficeResponse
	}{
		{
			name: "OK",
			req: &customer.CreateOfficeRequest{
				Name:    "Test name",
				Address: "Test address",
			},
			mockBehavior: func(s *mock_repository.MockOffice, expectedUUID uuid.UUID, expectedOffice model.Office) {
				s.EXPECT().CreateOffice(expectedUUID, expectedOffice).Return(nil)
			},
			expectedUUID: uuid.New(),
			expectedOffice: model.Office{
				Uuid:    uuid.New(),
				Name:    "Test name",
				Address: "Test Address",
			},
			expectedError:  nil,
			expectedResult: &customer.CreateOfficeResponse{},
		},
	}
	for _, testCase := range testTable {
		c := gomock.NewController(t)
		defer c.Finish()

		officeRepo := mock_repository.NewMockOffice(c)
		service := NewOfficeService(officeRepo)

		testCase.mockBehavior(officeRepo, testCase.expectedUUID, testCase.expectedOffice)

		result, err := service.CreateOffice(context.Background(), testCase.req)

		assert.Equal(t, testCase.expectedError, err)
		assert.Equal(t, testCase.expectedResult, err)

		officeRepo.
	}
}
