package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
)

type OfficeRepository struct {
	db *sqlx.DB
}

func NewOfficeRepository(db *sqlx.DB) *OfficeRepository {
	return &OfficeRepository{
		db: db,
	}
}

func (or *OfficeRepository) CreateOffice(uuId uuid.UUID, office model.Office) error {
	const query = `insert into offices (uuid, name, address) values ($1, $2, $3)`

	_, err := or.db.Exec(query, uuId, office.Name, office.Address)

	return err
}

func (or *OfficeRepository) GetOfficesList() ([]*model.Office, error) {
	const query = `select uuid, name, address, created_at from offices`

	var officeList []*model.Office

	err := or.db.Select(&officeList, query)

	return officeList, err
}

func (or *OfficeRepository) GetOffice(officeUuid uuid.UUID) (model.Office, error) {
	const query = `select * from customer.offices where uuid=$1`

	var office model.Office

	err := or.db.Get(&office, query, officeUuid)

	return office, err
}

func (or *OfficeRepository) GetOfficeByUserUUID(userUUID uuid.UUID) (model.Office, error) {
	const query = `select o.uuid, o.name, o.address
					from customer.offices o
							 join users u on o.uuid = u.office_uuid
					where u.uuid = $1`
	var office model.Office

	err := or.db.Get(&office, query, userUUID)

	return office, err
}
