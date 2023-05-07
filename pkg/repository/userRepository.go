package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/pkg/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(uuId uuid.UUID, user model.User) error {
	const query = `insert into customer.users(uuid, name, office_uuid, office_name) VALUES ($1, $2, $3, $4)`

	_, err := ur.db.Exec(query, uuId, user.Name, user.OfficeUuid, user.OfficeName)

	return err
}

func (ur *UserRepository) GetUsersList(officeUuid uuid.UUID) ([]*model.User, error) {
	const query = `select uuid, name, office_uuid, office_name, created_at from customer.users where office_uuid=$1`

	var usersList []*model.User

	err := ur.db.Select(&usersList, query, officeUuid)

	return usersList, err
}
