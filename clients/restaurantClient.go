package clients

import (
	"context"
	"fmt"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ServiceClient struct {
	OrderServiceClient   restaurant.OrderServiceClient
	MenuServiceClient    restaurant.MenuServiceClient
	ProductServiceClient restaurant.ProductServiceClient
}

type RestaurantServiceClient struct {
	Client ServiceClient
}

func InitRestaurantServiceClient(url string) RestaurantServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect to restaurantService:", err)
	}

	c := RestaurantServiceClient{
		Client: ServiceClient{
			OrderServiceClient:   restaurant.NewOrderServiceClient(cc),
			MenuServiceClient:    restaurant.NewMenuServiceClient(cc),
			ProductServiceClient: restaurant.NewProductServiceClient(cc),
		},
	}

	return c
}

func (c *RestaurantServiceClient) GetMenu(onDateMenu time.Time) (*restaurant.GetMenuResponse, error) {
	req := &restaurant.GetMenuRequest{OnDate: timestamppb.New(onDateMenu)}

	return c.Client.MenuServiceClient.GetMenu(context.Background(), req)
}
