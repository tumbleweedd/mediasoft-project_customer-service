package service

import (
	"context"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/repository"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
)

type OrderService struct {
	repository repository.Order
}

func NewOrderService(r repository.Order) *OrderService {
	return &OrderService{repository: r}
}

func (os *OrderService) CreateOrder(ctx context.Context, request *customer.CreateOrderRequest) (*customer.CreateOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (os *OrderService) GetActualMenu(ctx context.Context, request *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error) {
	//TODO implement me
	panic("implement me")
}
