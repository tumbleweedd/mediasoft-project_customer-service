package model

import (
	"github.com/google/uuid"
	"time"
)

type Office struct {
	Uuid    uuid.UUID `json:"uuid" db:"uuid"`
	Name    string    `json:"name" db:"name"`
	Address string    `json:"address" db:"address"`
	//CompanyUuid uuid.UUID `json:"company_uuid" db:"company_uuid"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
