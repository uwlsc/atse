package orchestrators

import (
	"magazine_api/api/serializers/requests"
	"magazine_api/lib"
	"magazine_api/models"
	"magazine_api/services"

	"github.com/google/uuid"
)

type EmployeeProfileOrchestrator struct {
	logger          lib.Logger
	userService     services.UserService
	cognito_service services.CognitoAuthService
}

func NewEmployeeProfileOrchestrator(
	l lib.Logger,
	u services.UserService,
	cognito_service services.CognitoAuthService,
) EmployeeProfileOrchestrator {
	return EmployeeProfileOrchestrator{
		logger:          l,
		userService:     u,
		cognito_service: cognito_service,
	}
}

func (e EmployeeProfileOrchestrator) CreateEmployee(user requests.CreateUser, id uuid.UUID,
) (*models.EmployeeProfile, error) {
	profile := &models.EmployeeProfile{
		EmployeeProfileBase: models.EmployeeProfileBase{
			UserId:        id,
			CompanyId:     user.CompanyId,
			Name:          user.Name,
			Role:          user.Role,
			Email:         user.Email,
			ContactNumber: user.ContactNumber,
			Picture:       user.Picture,
		},
	}

	return profile, nil
}
