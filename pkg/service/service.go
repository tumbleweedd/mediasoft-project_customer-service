package service

import (
	"context"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/rabbitmq"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/repository"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
)

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
	rabbitMQConn *rabbitmq.RabbitMQConn
}

func NewService(r *repository.Repository, rabbitMQConn *rabbitmq.RabbitMQConn) *Service {
	return &Service{
		Office:       NewOfficeService(r.Office),
		Order:        NewOrderService(r.Order, rabbitMQConn),
		User:         NewUserService(r.User, r.Office),
		rabbitMQConn: rabbitMQConn,
	}
}
