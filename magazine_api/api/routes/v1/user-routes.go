package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

// UserRoutes struct
type UserRoutes struct {
	logger           lib.Logger
	handler          infrastructure.Router
	pagination       middlewares.PaginationMiddleware
	uploadMiddleware middlewares.UploadMiddleware
	userController   handlers.UserHandler
}

func NewUserRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	uploadMiddleware middlewares.UploadMiddleware,
	userController handlers.UserHandler) UserRoutes {
	return UserRoutes{
		handler:          handler,
		logger:           logger,
		pagination:       pagination,
		userController:   userController,
		uploadMiddleware: uploadMiddleware,
	}
}

// Setup user routes
func (s UserRoutes) Setup(handler *gin.RouterGroup) {
	s.logger.Info("Setting up User routes")
	api := handler.Group("/user")
	{
		api.POST("", s.userController.CreateUser)
	}
}
