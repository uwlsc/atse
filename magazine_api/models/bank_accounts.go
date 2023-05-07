package models

import "github.com/google/uuid"

type AccountBase struct {
	CreatorId *uuid.UUID `json:"creator_id"`

	Name *string `json:"name"`

	AccountNumber *string `json:"account_number"`
	BankName      *string `json:"bank_name"`
	BankBranch    *string `json:"bank_branch"`
}

type BankAccount struct {
	AccountBase
	Base
	BaseDate
}
