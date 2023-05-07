package model

import "github.com/google/uuid"

type Company struct {
	Id         int       `json:"-" db:"id"`
	Name       string    `json:"name" db:"name"`
	OfficeUuid uuid.UUID `json:"office_uuid" db:"office_uuid"`
}
