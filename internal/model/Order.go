package model

import "github.com/google/uuid"

type Order struct {
	OrderUUID uuid.UUID `json:"order_uuid" db:"uuid"`
	UserUUID  uuid.UUID `json:"user_uuid" db:"user_uuid"`
}

type OrderItem struct {
	ID          int       `json:"id" db:"id"`
	Count       int       `json:"count" db:"count"`
	ProductUUID uuid.UUID `json:"product_uuid" db:"product_uuid"`
	OrderUUID   uuid.UUID `json:"order_uuid" db:"order_uuid"`
}

type OrderItemByOffice struct {
	Count       int       `json:"Count"`
	ProductUUID uuid.UUID `json:"ProductUUID"`
}

type OrdersByOffice struct {
	UserUUID      uuid.UUID            `json:"UserUUID"`
	OfficeUUID    uuid.UUID            `json:"Office_uuid"`
	OfficeName    string               `json:"Office_name"`
	OfficeAddress string               `json:"Office_address"`
	Salads        []*OrderItemByOffice `json:"Salads,omitempty"`
	Garnishes     []*OrderItemByOffice `json:"Garnishes,omitempty"`
	Meats         []*OrderItemByOffice `json:"Meats,omitempty"`
	Soups         []*OrderItemByOffice `json:"Soups,omitempty"`
	Drinks        []*OrderItemByOffice `json:"Drinks,omitempty"`
	Desserts      []*OrderItemByOffice `json:"Desserts,omitempty"`
}
