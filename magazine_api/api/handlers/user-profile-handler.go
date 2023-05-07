package handlers

import (
	"magazine_api/lib"
	"magazine_api/models"
	"magazine_api/services"
	"time"

	"github.com/danhper/structomap"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileHandler struct {
	logger  lib.Logger
	service services.UserProfileService
}

func NewUserProfileHandler(logger lib.Logger, service services.UserProfileService) UserProfileHandler {
	return UserProfileHandler{
		logger:  logger,
		service: service,
	}
}

// CreateUserProfile godoc
// @Summary      Create UserProfile
// @Description  It creates UserProfile
// @Tags         UserProfile
// @Accept       json
// @Produce      json
// @Param        profile  body      models.UserProfile  true  "Add UserProfile"
// @Success      200      {object}  object{data=models.UserProfile}
// @Router       /user_profile [post]
//
// Creates user profile for employee
func (s UserProfileHandler) CreateUserProfile(c *gin.Context) {
	var user_profile *models.UserProfile
	if err := c.ShouldBind(&user_profile); err != nil {
		handleError(s.logger, c, err)
		return
	}

	// Create user_profile for employee in our Database
	user_profile, err := s.service.CreateUserProfile(user_profile)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": user_profile})
}

// GetOneUserProfileByID godoc
// @Summary      Gets One User Profile by ID
// @Description  Gets One user profile by ID
// @Tags         UserProfile
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  models.UserProfile
// @Router       /user_profile/id/{id} [get]
// Gets Users Profile By ID controller
func (u UserProfileHandler) GetUserProfileByID(c *gin.Context) {
	id := c.Param("id")

	user, err := u.service.GetUserProfileByID(uuid.MustParse(id))
	if err != nil {
		handleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": user})
}

// UpdateUserProfile godoc
// @Summary      Update UserProfile
// @Description  Updates UserProfile of user
// @Tags         UserProfile
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Update user_profile"
// @Success      200  {object}  object{data=models.UserProfile}
// @Router       /user_profile/{id} [patch]
// Patch user_profile of creator by Id controller
func (a UserProfileHandler) PatchUserProfile(c *gin.Context) {
	id := c.Param("id")

	user_profile, err := a.service.GetUserProfileByID(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	var newUserProfile models.UserProfileBase
	if err := c.ShouldBindJSON(&newUserProfile); err != nil {
		handleError(a.logger, c, err)
		return
	}

	user_profileMap := structomap.New().UseSnakeCase().PickAll().
		Omit("UserId").
		OmitIf(func(ch interface{}) bool {
			return newUserProfile.Name == nil
		}, "Name").
		OmitIf(func(ch interface{}) bool {
			return newUserProfile.Email == nil
		}, "Email").
		OmitIf(func(ch interface{}) bool {
			return newUserProfile.ContactNumber == nil
		}, "ContactNumber").
		OmitIf(func(ch interface{}) bool {
			return newUserProfile.Picture == nil
		}, "Picture").
		Transform(newUserProfile)

	if len(user_profileMap) > 0 {
		user_profileMap["updated_on"] = time.Now()
		user_profileMap["id"] = user_profile.ID

		err := a.service.UpdateUserProfile(user_profile.ID, &user_profileMap)
		if err != nil {
			handleError(a.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": user_profileMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}

// DeleteUserProfile godoc
// @Summary      Soft Delete an user_profile
// @Description  Delete by user_profile ID
// @Tags         UserProfile
// @Produce      json
// @Param        id   path      string  true  "Unique ID"
// @Success      204  {object}  object{data=string}
// @Router       /user_profile/{id} [delete]
//
// Delete UserProfile By ID controller
func (a UserProfileHandler) DeleteUserProfileByID(c *gin.Context) {
	id := c.Param("id")

	err := a.service.DeleteUserProfile(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}
