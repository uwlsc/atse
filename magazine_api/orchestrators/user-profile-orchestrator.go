package orchestrators

import (
	"magazine_api/api/serializers/requests"
	"magazine_api/lib"
	"magazine_api/models"
	"magazine_api/services"

	"github.com/google/uuid"
)

type UserProfileOrchestrator struct {
	logger         lib.Logger
	userService    services.UserService
	profileService services.UserProfileService
}

func NewUserProfileOrchestrator(
	logger lib.Logger,
	userService services.UserService,
	profileService services.UserProfileService,
) UserProfileOrchestrator {
	return UserProfileOrchestrator{
		logger:         logger,
		userService:    userService,
		profileService: profileService,
	}
}

func (e UserProfileOrchestrator) CreateUserProfile(
	user requests.CreateUser,
	id uuid.UUID,
) (*models.UserProfile, error) {
	profile := &models.UserProfile{
		UserProfileBase: models.UserProfileBase{
			UserId:        id,
			Name:          user.Name,
			Email:         user.Email,
			ContactNumber: user.ContactNumber,
			Picture:       user.Picture,
		},
	}

	profile, err := e.profileService.CreateUserProfile(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
