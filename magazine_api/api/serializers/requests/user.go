package requests

import (
	"magazine_api/lib"
	"magazine_api/models"
)

type CreateUser struct {
	models.UserBase
	Picture          *lib.SignedURL        `json:"picture"`
	Address          *models.AddressBase   `json:"address"`
	EmergencyContact *models.ContactBase   `json:"emergency_contact"`
	CompanyId        *string               `json:"company_id"`
	Salary           *models.SalaryBase    `json:"salary"`
	BankAccount      []models.AccountBase  `json:"bank_account"`
	Documents        []models.DocumentBase `json:"documents"`
}
