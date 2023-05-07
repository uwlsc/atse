package models

import "github.com/google/uuid"

type AddressBase struct {
	CreatorId     uuid.UUID `json:"creator_id"`
	Street        *string   `json:"street"`
	Ward          *int      `json:"ward"`
	Municipality  *string   `json:"municipality"`
	District      *string   `json:"district"`
	State         *string   `json:"state"`
	Country       *string   `json:"country"`
	ContactNumber []*string `json:"contact_number"`
}

type Address struct {
	AddressBase
	Base
	BaseDate
}
