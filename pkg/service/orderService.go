package service

import (
	"context"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/clients"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/repository"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
)

type OrderService struct {
	orderRepository   repository.Order
	restaurantService clients.RestaurantServiceClient
}

func NewOrderService(r repository.Order) *OrderService {
	return &OrderService{orderRepository: r}
}

func (os *OrderService) CreateOrder(ctx context.Context, request *customer.CreateOrderRequest) (*customer.CreateOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (os *OrderService) GetActualMenu(ctx context.Context, request *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error) {
	//menu, err := os.restaurantService.GetMenu(time.Now())
	//if err != nil {
	//	return nil, status.Error(codes.Internal, fmt.Sprintf("restaurant service: %v", err.Error()))
	//}

	return &customer.GetActualMenuResponse{}, nil
}
