package service

import (
	"context"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/clients"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/model"
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
	order := &model.Order{
		UserUUID:  request.UserUuid,
		Salads:    make([]model.OrderProduct, len(request.Salads)),
		Garnishes: make([]model.OrderProduct, len(request.Garnishes)),
		Meats:     make([]model.OrderProduct, len(request.Meats)),
		Soups:     make([]model.OrderProduct, len(request.Soups)),
		Drinks:    make([]model.OrderProduct, len(request.Drinks)),
		Desserts:  make([]model.OrderProduct, len(request.Desserts)),
	}

	for i, p := range request.Salads {
		order.Salads[i] = model.OrderProduct{Count: int(p.Count), ProductUUID: p.ProductUuid}
	}

	for i, p := range request.Garnishes {
		order.Garnishes[i] = model.OrderProduct{Count: int(p.Count), ProductUUID: p.ProductUuid}
	}

	for i, p := range request.Meats {
		order.Meats[i] = model.OrderProduct{Count: int(p.Count), ProductUUID: p.ProductUuid}
	}

	for i, p := range request.Soups {
		order.Soups[i] = model.OrderProduct{Count: int(p.Count), ProductUUID: p.ProductUuid}
	}

	for i, p := range request.Drinks {
		order.Drinks[i] = model.OrderProduct{Count: int(p.Count), ProductUUID: p.ProductUuid}
	}

	err := os.orderRepository.

}

func (os *OrderService) GetActualMenu(ctx context.Context, request *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error) {
	//menu, err := os.restaurantService.GetMenu(time.Now())
	//if err != nil {
	//	return nil, status.Error(codes.Internal, fmt.Sprintf("restaurant service: %v", err.Error()))
	//}

	return &customer.GetActualMenuResponse{}, nil
}

func formatOrder(orderItem []*customer.OrderItem, order *model.Order) {
	for i, p := range orderItem {
		order.Salads[i] = model.OrderProduct{Count: int(p.Count), ProductUUID: p.ProductUuid}
	}
}
