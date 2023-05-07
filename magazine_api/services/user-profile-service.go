package services

import (
	"magazine_api/component"
	"magazine_api/lib"
	"magazine_api/models"
	"time"

	"github.com/google/uuid"
)

type UserProfileService struct {
	logger lib.Logger
	comp   component.UserProfileComponent
}

func NewUserProfileService(logger lib.Logger, comp component.UserProfileComponent) UserProfileService {
	return UserProfileService{logger: logger, comp: comp}
}

func (u UserProfileService) CreateUserProfile(user *models.UserProfile) (*models.UserProfile, error) {
	user = u.BeforeCreate(user)

	err := u.comp.CreateUserProfile(*user)

	if err != nil {
		return nil, err
	}

	return user, nil

}

// Gets Cutting by ID from database
func (u UserProfileService) GetUserProfileByID(id uuid.UUID) (*models.UserProfile, error) {
	profile, err := u.comp.GetUserProfileFromID(id)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

// Update UserProfilement by Id in our database
func (u UserProfileService) UpdateUserProfile(id uuid.UUID, patch *map[string]interface{}) error {
	err := u.comp.PatchUserProfile(id, patch)
	if err != nil {
		return err
	}

	return nil
}

// Delete cutting assign by ID in our database
func (u UserProfileService) DeleteUserProfile(id uuid.UUID) error {
	err := u.comp.DeleteUserProfile(id)
	if err != nil {
		return err
	}

	return nil
}

// Permanent Delete user by in our database permanently
func (u UserProfileService) PermanentDeleteUserProfile(id uuid.UUID) error {
	err := u.comp.PermanentDeleteUserProfile(id)
	if err != nil {
		return err
	}

	return nil
}

func (u UserProfileService) BeforeCreate(profile *models.UserProfile) *models.UserProfile {
	profile.ID = uuid.New()
	create := time.Now()
	profile.CreatedOn = &create
	profile.UpdatedOn = &create

	return profile
}
