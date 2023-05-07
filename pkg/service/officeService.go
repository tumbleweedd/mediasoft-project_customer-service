package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/model"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/repository"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OfficeService struct {
	repository repository.Office
}

func NewOfficeService(r repository.Office) *OfficeService {
	return &OfficeService{repository: r}
}

func (os *OfficeService) CreateOffice(ctx context.Context, req *customer.CreateOfficeRequest) (*customer.CreateOfficeResponse, error) {
	office := model.Office{Name: req.Name, Address: req.Address}
	uuId := uuid.New()

	err := os.repository.CreateOffice(uuId, office)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &customer.CreateOfficeResponse{}, nil
}

func (os *OfficeService) GetOfficeList(ctx context.Context, req *customer.GetOfficeListRequest) (*customer.GetOfficeListResponse, error) {
	officeList, err := os.repository.GetOfficesList()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var data []*customer.Office

	for _, office := range officeList {
		data = append(data, &customer.Office{
			Uuid:      office.Uuid.String(),
			Name:      office.Name,
			Address:   office.Address,
			CreatedAt: timestamppb.New(office.CreatedAt),
		})
	}

	return &customer.GetOfficeListResponse{
		Result: data,
	}, nil
}
