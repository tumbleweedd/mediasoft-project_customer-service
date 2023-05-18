package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) CreateOrder(order model.Order, salads, garnishes, meats, soups, drinks, desserts []*customer.OrderItem) error {
	tx, err := or.db.Beginx()
	if err != nil {
		return err
	}

	var orderUUID uuid.UUID
	const createOrderQuery = `insert into customer.orders (uuid, user_uuid) VALUES ($1, $2) returning uuid`
	row := tx.QueryRow(createOrderQuery, order.OrderUUID, order.UserUUID)
	if err := row.Scan(&orderUUID); err != nil {
		tx.Rollback()
		return err
	}

	const createItemQuery = `insert into customer.order_items (count, product_uuid, order_uuid)
								values ($1, $2, $3)`
	for _, items := range [][]*customer.OrderItem{salads, garnishes, meats, soups, drinks, desserts} {
		for _, item := range items {
			_, err := tx.Exec(createItemQuery, item.Count, item.ProductUuid, orderUUID)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit()
}
