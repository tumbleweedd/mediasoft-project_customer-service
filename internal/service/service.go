package service

import (
	"context"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/repository"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/broker/kafka/producer"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Office interface {
	CreateOffice(context.Context, *customer.CreateOfficeRequest) (*customer.CreateOfficeResponse, error)
	GetOfficeList(context.Context, *customer.GetOfficeListRequest) (*customer.GetOfficeListResponse, error)
}

type Order interface {
	CreateOrder(context.Context, *customer.CreateOrderRequest) (*customer.CreateOrderResponse, error)
	GetActualMenu(context.Context, *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error)
}

type User interface {
	CreateUser(context.Context, *customer.CreateUserRequest) (*customer.CreateUserResponse, error)
	GetUserList(context.Context, *customer.GetUserListRequest) (*customer.GetUserListResponse, error)
}

type Service struct {
	Office
	User
	Order
	customer.UnsafeOfficeServiceServer
	customer.UnsafeOrderServiceServer
	customer.UnsafeUserServiceServer
	kafkaProducer *producer.Producer
	Done          chan struct{}
}

func NewService(r *repository.Repository, kafkaProducer *producer.Producer, Done chan struct{}) *Service {
	return &Service{
		Office:        NewOfficeService(r.Office),
		Order:         NewOrderService(r.Order, r.Office, kafkaProducer, Done),
		User:          NewUserService(r.User, r.Office),
		kafkaProducer: kafkaProducer,
	}
}
