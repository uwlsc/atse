package services

import (
	"magazine_api/api/serializers/responses"
	"magazine_api/component"
	"magazine_api/constants"
	"magazine_api/lib"
	"magazine_api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserService service layer
type UserService struct {
	logger lib.Logger
	repo   component.UserComponent
}

// NewUserService creates new instance of UserService
func NewUserService(logger lib.Logger, repo component.UserComponent) UserService {
	return UserService{logger: logger, repo: repo}
}

// CreateUser Creates the user in database
func (u UserService) CreateUser(user *models.User) (*models.User, error) {
	user = u.BeforeCreate(user)

	err := u.repo.CreateUser(*user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ListsUsers Lists users from database
func (u UserService) ListEmployees(c *gin.Context) (gin.H, error) {

	users, err := u.repo.ListEmployees(c.MustGet(constants.Limit).(int64), c.MustGet(constants.Offset).(int64))
	if err != nil {
		return nil, err
	}

	limit := c.MustGet(constants.Limit).(int64)
	size := c.MustGet(constants.Page).(int64)

	return gin.H{
		"data":       users["data"],
		"pagination": gin.H{"has_next": (users["count"].(int64) - limit*size) > 0, "count": users["count"]},
	}, nil
}

// ListEmployeesByType Lists users by type from database
func (u UserService) ListEmployeesByType(c *gin.Context, userType string) ([]*responses.EmployeeSmall, error) {
	users, err := u.repo.ListEmployeesByType(
		c.MustGet(constants.Limit).(int64), c.MustGet(constants.Page).(int64), userType,
	)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// ListsDeletedUsers Lists Deleted users from database
func (u UserService) ListsDeletedUsers() ([]*models.User, error) {
	users, err := u.repo.ListDeletedUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID Gets user by id from database
func (u UserService) GetProfileByID(id uuid.UUID) (*responses.EmployeeAll, error) {
	user, err := u.repo.GetProfileFromID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail Gets user by email from database
func (u UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.repo.GetUserFromEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByContactNumber Gets user by contact_number from database
func (u UserService) GetUserByContactNumber(contact_number string) (*models.User, error) {
	user, err := u.repo.GetUserFromContactNumber(contact_number)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID Gets user by id from database
func (u UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := u.repo.GetUserFromID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser Update user by in our database
func (u UserService) UpdateUser(id uuid.UUID, patch *map[string]interface{}) error {
	err := u.repo.PatchUser(id, patch)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser Delete user by in our database
func (u UserService) DeleteUser(id uuid.UUID) error {
	err := u.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

// PermanentDeleteUser deletes user permanently from database
func (u UserService) PermanentDeleteUser(id uuid.UUID) error {
	err := u.repo.PermanentDeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAllProfilesOfUser deletes user permanently from database
func (u UserService) DeleteAllProfilesOfUser(id uuid.UUID) error {
	err := u.repo.PermanentDeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (u UserService) BeforeCreate(user *models.User) *models.User {
	user.ID = uuid.New()
	create := time.Now()
	user.CreatedOn = &create
	user.UpdatedOn = &create

	return user
}
