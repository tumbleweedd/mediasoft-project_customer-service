package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/clients"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/model"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/rabbitmq"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/repository"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
	orderRepository   repository.Order
	restaurantService clients.RestaurantServiceClient
	rabbitMQConn      *rabbitmq.RabbitMQConn
}

func NewOrderService(r repository.Order, rabbitMQConn *rabbitmq.RabbitMQConn) *OrderService {
	return &OrderService{
		orderRepository: r,
		rabbitMQConn:    rabbitMQConn,
	}
}

func (os *OrderService) CreateOrder(ctx context.Context, request *customer.CreateOrderRequest) (*customer.CreateOrderResponse, error) {
	userUUID, err := uuid.Parse(request.UserUuid)
	if err != nil {
		return nil, err
	}

	order := model.Order{
		OrderUUID: uuid.New(),
		UserUUID:  userUUID,
	}

	if err := os.orderRepository.CreateOrder(
		order,
		request.Salads,
		request.Garnishes,
		request.Meats,
		request.Soups,
		request.Drinks,
		request.Desserts); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := os.rabbitMQConn.SendOrder(&order); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &customer.CreateOrderResponse{}, err
}

func (os *OrderService) GetActualMenu(ctx context.Context, request *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error) {
	//menu, err := os.restaurantService.GetMenu(time.Now())
	//if err != nil {
	//	return nil, status.Error(codes.Internal, fmt.Sprintf("restaurant service: %v", err.Error()))
	//}

	return &customer.GetActualMenuResponse{}, nil
}
