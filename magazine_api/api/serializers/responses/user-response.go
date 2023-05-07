package responses

import (
	"magazine_api/lib"
	"magazine_api/models"
	"time"

	"github.com/google/uuid"
)

type EmployeeSmall struct {
	Name          *string         `json:"name"`
	Email         *string         `json:"email"`
	Role          models.UserRole `json:"role"`
	ContactNumber *string         `json:"contact_number"`

	UserId *uuid.UUID `json:"user_id"`

	EmployeeProfileId *uuid.UUID `json:"employee_profile_id"`
	CompanyId         *string    `json:"company_id"`
}

type EmployeeAll struct {
	EmployeeSmall

	Picture *lib.SignedURL `json:"picture"`

	Salary []*struct {
		Amount        *float32             `json:"amount"`
		SalaryFormat  *models.SalaryFormat `json:"salary_format"`
		TailorRate    []models.TailorRate  `json:"tailor_rate"`
		EffectiveFrom *models.Month        `json:"effective_from"`
		EffectiveTo   *models.Month        `json:"effective_to"`
	} `json:"salary"`

	BankAccount []*models.AccountBase `json:"bank_account"`

	Address []*models.AddressBase `json:"address"`

	EmergencyContact []*models.ContactBase `json:"emergency_contact"`

	Document []*models.DocumentBase `json:"document"`

	DeletedOn *time.Time `json:"deleted_on" form:"deleted_on"`
	CreatedOn *time.Time `json:"created_on" form:"created_on"`
}
