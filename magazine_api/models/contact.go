package models

import "github.com/google/uuid"

type ContactBase struct {
	CreatorId uuid.UUID `json:"creator_id"`

	Name          *string `json:"name"`
	Relation      *string `json:"relation"`
	ContactNumber *string `json:"contact_number"`
}

type Contact struct {
	ContactBase
	Base
	BaseDate
}
