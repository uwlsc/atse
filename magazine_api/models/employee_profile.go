package models

import (
	"magazine_api/lib"

	"github.com/google/uuid"
)

// EmployeeProfileBase Employee profile
type EmployeeProfileBase struct {
	// Reference to the main user
	UserId uuid.UUID `json:"user_id"`

	CompanyId     *string        `json:"company_id"`
	Name          *string        `json:"name"`
	Role          UserRole       `json:"role"`
	Email         *string        `json:"email"`
	ContactNumber *string        `json:"contact_number"`
	Picture       *lib.SignedURL `json:"image"`
}

type EmployeeProfile struct {
	Base
	BaseDate
	EmployeeProfileBase
}
