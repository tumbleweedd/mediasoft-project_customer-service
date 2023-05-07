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

type UserService struct {
	userRepo   repository.User
	officeRepo repository.Office
}

func NewUserService(userRepo repository.User, officeRepo repository.Office) *UserService {
	return &UserService{userRepo: userRepo, officeRepo: officeRepo}
}
func (us *UserService) CreateUser(ctx context.Context, request *customer.CreateUserRequest) (*customer.CreateUserResponse, error) {

	officeUuid, err := uuid.Parse(request.OfficeUuid)
	if err != nil {
		return nil, err
	}

	office, err := us.officeRepo.GetOffice(officeUuid)
	if err != nil {
		return nil, err
	}

	user := model.User{Name: request.Name, OfficeUuid: officeUuid, OfficeName: office.Name}

	userUuid := uuid.New()

	err = us.userRepo.CreateUser(userUuid, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &customer.CreateUserResponse{}, nil
}

func (us *UserService) GetUserList(ctx context.Context, request *customer.GetUserListRequest) (*customer.GetUserListResponse, error) {
	officeUuid, err := uuid.Parse(request.OfficeUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	users, err := us.userRepo.GetUsersList(officeUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var data []*customer.User
	for _, user := range users {
		data = append(data, &customer.User{
			Uuid:       user.Uuid.String(),
			Name:       user.Name,
			OfficeUuid: user.OfficeUuid.String(),
			OfficeName: user.OfficeName,
			CreatedAt:  timestamppb.New(user.CreatedAt),
		})
	}

	return &customer.GetUserListResponse{Result: data}, nil
}
