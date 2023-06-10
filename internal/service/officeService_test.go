package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
	mock_repository "github.com/tumbleweedd/mediasoft-intership/customer-service/internal/repository/mocks"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestOfficeService_CreateOffice(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockOffice, req *customer.CreateOfficeRequest)

	testTable := []struct {
		name           string
		req            *customer.CreateOfficeRequest
		expectedOffice model.Office
		mockBehavior   mockBehavior
		wantError      bool
		expectedResult *customer.CreateOfficeResponse
	}{
		{
			name: "OK",
			req: &customer.CreateOfficeRequest{
				Name:    "Test name",
				Address: "Test address",
			},
			mockBehavior: func(mockRepo *mock_repository.MockOffice, req *customer.CreateOfficeRequest) {
				expOffice := model.Office{Name: req.Name, Address: req.Address}
				mockRepo.EXPECT().CreateOffice(gomock.Any(), expOffice).Return(nil)
			},
			wantError:      false,
			expectedResult: &customer.CreateOfficeResponse{},
		},
		{
			name: "Error creating office",
			req: &customer.CreateOfficeRequest{
				Name:    "Test name",
				Address: "Test address",
			},
			mockBehavior: func(mockRepo *mock_repository.MockOffice, req *customer.CreateOfficeRequest) {
				expOffice := model.Office{Name: req.Name, Address: req.Address}
				mockRepo.EXPECT().CreateOffice(gomock.Any(), expOffice).Return(errors.New("some error"))
			},
			wantError:      true,
			expectedResult: nil,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			officeRepo := mock_repository.NewMockOffice(c)
			service := NewOfficeService(officeRepo)

			testCase.mockBehavior(officeRepo, testCase.req)

			result, err := service.CreateOffice(context.Background(), testCase.req)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedResult, result)
			}
		})
	}
}

func TestOfficeService_GetOfficeList(t *testing.T) {
	expectedOffices := []*model.Office{
		{
			Uuid:      uuid.New(),
			Name:      "Office 1",
			Address:   "Address 1",
			CreatedAt: time.Now(),
		},
		{
			Uuid:      uuid.New(),
			Name:      "Office 2",
			Address:   "Address 2",
			CreatedAt: time.Now(),
		},
	}

	type mockBehavior func(mockRepo *mock_repository.MockOffice, expectedOffices []*model.Office)

	testTable := []struct {
		name            string
		expectedOffices []*model.Office
		mockBehavior    mockBehavior
		wantError       bool
		expectedResult  *customer.GetOfficeListResponse
	}{
		{
			name:            "OK",
			expectedOffices: expectedOffices,
			mockBehavior: func(mockRepo *mock_repository.MockOffice, expectedOffices []*model.Office) {
				mockRepo.EXPECT().GetOfficesList().Return(expectedOffices, nil)
			},
			wantError: false,
			expectedResult: &customer.GetOfficeListResponse{
				Result: []*customer.Office{
					{
						Uuid:      expectedOffices[0].Uuid.String(),
						Name:      expectedOffices[0].Name,
						Address:   expectedOffices[0].Address,
						CreatedAt: timestamppb.New(expectedOffices[0].CreatedAt),
					},
					{
						Uuid:      expectedOffices[1].Uuid.String(),
						Name:      expectedOffices[1].Name,
						Address:   expectedOffices[1].Address,
						CreatedAt: timestamppb.New(expectedOffices[1].CreatedAt),
					},
				},
			},
		},
		{
			name:            "No offices",
			expectedOffices: []*model.Office{},
			mockBehavior: func(mockRepo *mock_repository.MockOffice, expectedOffices []*model.Office) {
				mockRepo.EXPECT().GetOfficesList().Return(expectedOffices, nil)
			},
			wantError:      false,
			expectedResult: &customer.GetOfficeListResponse{Result: []*customer.Office(nil)},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			officeRepo := mock_repository.NewMockOffice(c)
			service := NewOfficeService(officeRepo)

			testCase.mockBehavior(officeRepo, testCase.expectedOffices)

			result, err := service.GetOfficeList(context.Background(), &customer.GetOfficeListRequest{})
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedResult, result)
			}
		})
	}
}
