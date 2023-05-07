package models

import "github.com/google/uuid"

type Month int

const (
	Shrawan Month = iota + 1
	Bhadra
	Ashoj
	Karthik
	Manghsir
	Poush
	Magh
	Falgun
	Chaitra
	Baisakh
	Jyestha
	Ashad
)

type SalarySheetBase struct {
	EmployeeId   *uuid.UUID
	EmployeeName *string
	CompanyId    *string

	// Month
	Month Month

	// Addition
	BaseSalary  *float32
	Bonus       *float32
	GrossSalary *float32 // BaseSalary + Bonus

	// Deduction
	TDS *float32

	// Deduction Amount
	TotalDeduction *float32

	// Net Earn
	NetEarn *float32

	// Advance Payments
	// Find out Advance Payments till now for this year.
	AdvancePayments *float32

	// Net Salary
	NetSalary *float32 // GrossSalary - TotalDeduction

	// Transaction Id
	TransactionId []*uuid.UUID `json:"transaction_id"`
}

type SalarySheet struct {
	SalarySheetBase
	Base
	BaseDate
}
