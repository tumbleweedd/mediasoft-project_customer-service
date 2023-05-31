package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	model2 "github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Office interface {
	CreateOffice(uuId uuid.UUID, office model2.Office) error
	GetOfficesList() ([]*model2.Office, error)
	GetOffice(officeUuid uuid.UUID) (model2.Office, error)
	GetOfficeByUserUUID(userUUID uuid.UUID) (model2.Office, error)
}

type Order interface {
	CreateOrder(order model2.Order, salads, garnishes, meats, soups, drinks, desserts []*customer.OrderItem) error
}

type User interface {
	CreateUser(uuId uuid.UUID, user model2.User) error
	GetUsersList(officeUuid uuid.UUID) ([]*model2.User, error)
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
