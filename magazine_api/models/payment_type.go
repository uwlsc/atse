package models

type PaidMedium int

const (
	Cash PaidMedium = iota + 1
	Bank
	Online
)

type PaidStatus int

const (
	Complete PaidStatus = iota + 1
	Pending
)

type PaymentTypeBase struct {
	PaymentName *string `json:"payment_name"`
}

type PaymentType struct {
	Base
	BaseDate
	BaseCreatedBy
	PaymentTypeBase
}
