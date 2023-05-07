package orchestrators

import (
	"magazine_api/api/serializers/requests"
	"magazine_api/lib"
	"magazine_api/models"
	"magazine_api/services"
)

type UserOrchestrator struct {
	logger          lib.Logger
	userService     services.UserService
	cognito_service services.CognitoAuthService
	employeeService EmployeeProfileOrchestrator
	profileService  UserProfileOrchestrator
}

func NewUserOrchestrator(
	l lib.Logger, u services.UserService, e EmployeeProfileOrchestrator, cognito_service services.CognitoAuthService,
	userProfile UserProfileOrchestrator,
) UserOrchestrator {
	return UserOrchestrator{
		logger:          l,
		userService:     u,
		employeeService: e,
		cognito_service: cognito_service,
		profileService:  userProfile,
	}
}

func (e UserOrchestrator) CreateUser(user_request requests.CreateUser) (*requests.CreateUser, error) {

	user := &models.User{
		UserBase: models.UserBase{
			Name:          user_request.Name,
			Password:      user_request.Password,
			Role:          user_request.Role,
			Email:         user_request.Email,
			ContactNumber: user_request.ContactNumber,
		},
	}

	// Create User in our Database
	user, err := e.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	id := user.ID.String()
	_, err = e.cognito_service.AdminCreateUser(&id, user.Email, user.ContactNumber, user.Password)
	if err != nil {
		// Delete user from our database upon error
		e.userService.PermanentDeleteUser(user.ID)

		return nil, err
	}

	if e.contains("employee", user.Role) {
		emp, err := e.employeeService.CreateEmployee(user_request, user.ID)
		if err != nil {
			// Delete user from our database upon error
			e.userService.PermanentDeleteUser(user.ID)

			return nil, err
		}
		e.cognito_service.SetEmployeeRoleToUser(&id, emp.ID.String())
	}

	if e.contains("user", user.Role) {
		use, err := e.profileService.CreateUserProfile(user_request, user.ID)
		if err != nil {
			// Delete user from our database upon error
			e.userService.PermanentDeleteUser(user.ID)

			return nil, err
		}

		e.cognito_service.SetUserRoleToUser(&id, use.ID.String())
	}

	e.cognito_service.SetRoleToUser(&id, user.Role)

	return &user_request, nil
}

func (e UserOrchestrator) contains(value string, checkIn []string) bool {
	for _, val := range checkIn {
		if value == val {
			return true
		}
	}

	return false
}
