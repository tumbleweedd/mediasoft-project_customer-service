package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Uuid       uuid.UUID `json:"uuid" db:"uuid"`
	Name       string    `json:"name" db:"name"`
	OfficeUuid uuid.UUID `json:"office_uuid" db:"office_uuid"`
	OfficeName string    `json:"office_name" db:"office_name"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
