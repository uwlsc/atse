package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

// UserProfileRoutes struct
type UserProfileRoutes struct {
	logger           lib.Logger
	handler          infrastructure.Router
	pagination       middlewares.PaginationMiddleware
	uploadMiddleware middlewares.UploadMiddleware
	userController   handlers.UserProfileHandler
}

func NewUserProfileRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	uploadMiddleware middlewares.UploadMiddleware,
	userController handlers.UserProfileHandler) UserProfileRoutes {
	return UserProfileRoutes{
		handler:          handler,
		logger:           logger,
		pagination:       pagination,
		userController:   userController,
		uploadMiddleware: uploadMiddleware,
	}
}

// Setup user routes
func (s UserProfileRoutes) Setup(handler *gin.RouterGroup) {
	s.logger.Info("Setting up UserProfile routes")
	api := handler.Group("/user_profile")
	{
		api.POST("", s.uploadMiddleware.Push(
			s.uploadMiddleware.Config().
				WebpEnable(true).
				ThumbEnable(true).
				Folder("profile_image")).
			Handle(), s.userController.CreateUserProfile)

		api.GET("/id/:id", s.userController.GetUserProfileByID)
		api.PATCH("/:id", s.userController.PatchUserProfile)
		api.DELETE("/:id", s.userController.DeleteUserProfileByID)
	}
}
