package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/model"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
)

type Office interface {
	CreateOffice(uuId uuid.UUID, office model.Office) error
	GetOfficesList() ([]*model.Office, error)
	GetOffice(officeUuid uuid.UUID) (model.Office, error)
}

type Order interface {
	CreateOrder(order model.Order, salads, garnishes, meats, soups, drinks, desserts []*customer.OrderItem) error
}

type User interface {
	CreateUser(uuId uuid.UUID, user model.User) error
	GetUsersList(officeUuid uuid.UUID) ([]*model.User, error)
}

type Repository struct {
	Office
	Order
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Office: NewOfficeRepository(db),
		Order:  NewOrderRepository(db),
		User:   NewUserRepository(db),
	}
}
