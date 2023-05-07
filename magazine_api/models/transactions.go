package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionBase struct {
	Title *string `json:"title"`

	TransactionCost *float32 `json:"transaction_cost"`

	DebitAmount  *float32 `json:"debit_amount"`
	CreditAmount *float32 `json:"credit_amount"`

	PaymentTo   *uuid.UUID `json:"payment_to"`
	PaymentFrom *uuid.UUID `json:"payment_from"`

	PaymentDate  *time.Time `json:"payment_date" form:"payment_date" time_format:"2006-01-02"`
	PaymentMonth *Month     `json:"payment_month"`

	PaidType   *uuid.UUID  `json:"payment_type"`
	PaidMedium *PaidMedium `json:"paid_medium"`

	BankPaymentFrom          *uuid.UUID `json:"bank_payment_from"`
	BankPaymentTo            *uuid.UUID `json:"bank_payment_to"`
	BankPaymentTransactionId *string    `json:"bank_payment_transaction_id"`

	OnlinePaymentName          *string `json:"online_payment_name"`
	OnlinePaymentFrom          *string `json:"online_payment_from"`
	OnlinePaymentTo            *string `json:"online_payment_to"`
	OnlinePaymentTransactionId *string `json:"online_payment_transaction_id"`

	Remarks *string `json:"remarks"`
}

type Transaction struct {
	Base
	BaseDate
	BaseCreatedBy
	TransactionBase
}
