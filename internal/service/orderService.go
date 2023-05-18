package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/clients"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/repository"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/broker/kafka"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
	orderRepository   repository.Order
	officeRepository  repository.Office
	restaurantService clients.RestaurantServiceClient
	kafkaProducer     *kafka.Producer
	Done              chan struct{}
}

func NewOrderService(orderRepo repository.Order, officeRepo repository.Office, kafkaProducer *kafka.Producer, Done chan struct{}) *OrderService {
	return &OrderService{
		orderRepository:  orderRepo,
		officeRepository: officeRepo,
		kafkaProducer:    kafkaProducer,
		Done:             Done,
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

	office, err := os.officeRepository.GetOfficeByUserUUID(order.UserUUID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	orderByOffice, err := buildOrderByOffice(&order, &office, request)

	go os.kafkaProducer.StartProduce(os.Done, "orders", *orderByOffice)

	return &customer.CreateOrderResponse{}, err
}

func (os *OrderService) GetActualMenu(ctx context.Context, request *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error) {
	//menu, err := os.restaurantService.GetMenu(time.Now())
	//if err != nil {
	//	return nil, status.Error(codes.Internal, fmt.Sprintf("restaurant service: %v", err.Error()))
	//}

	return &customer.GetActualMenuResponse{}, nil
}

func buildOrderByOffice(order *model.Order, office *model.Office, request *customer.CreateOrderRequest) (*model.OrdersByOffice, error) {
	orderByOffice := &model.OrdersByOffice{
		UserUUID:      order.UserUUID,
		OfficeUUID:    office.Uuid,
		OfficeName:    office.Name,
		OfficeAddress: office.Address,
	}

	if err := addOrderItems(&orderByOffice.Salads, request.Salads); err != nil {
		return nil, err
	}
	if err := addOrderItems(&orderByOffice.Garnishes, request.Garnishes); err != nil {
		return nil, err
	}
	if err := addOrderItems(&orderByOffice.Meats, request.Meats); err != nil {
		return nil, err
	}
	if err := addOrderItems(&orderByOffice.Soups, request.Soups); err != nil {
		return nil, err
	}
	if err := addOrderItems(&orderByOffice.Drinks, request.Drinks); err != nil {
		return nil, err
	}
	if err := addOrderItems(&orderByOffice.Desserts, request.Desserts); err != nil {
		return nil, err
	}

	return orderByOffice, nil
}

func addOrderItems(dest *[]*model.OrderItemByOffice, src []*customer.OrderItem) error {
	for _, item := range src {
		productUUID, err := uuid.Parse(item.ProductUuid)
		if err != nil {
			return err
		}

		*dest = append(*dest, &model.OrderItemByOffice{
			Count:       int(item.Count),
			ProductUUID: productUUID,
		})
	}

	return nil
}
