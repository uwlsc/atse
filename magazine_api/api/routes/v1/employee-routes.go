package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

// Employee struct
type EmployeeRoutes struct {
	logger           lib.Logger
	handler          infrastructure.Router
	pagination       middlewares.PaginationMiddleware
	uploadMiddleware middlewares.UploadMiddleware
	userController   handlers.EmployeeHandler
}

func NewEmployeeRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	uploadMiddleware middlewares.UploadMiddleware,
	userController handlers.EmployeeHandler) EmployeeRoutes {
	return EmployeeRoutes{
		handler:          handler,
		logger:           logger,
		pagination:       pagination,
		userController:   userController,
		uploadMiddleware: uploadMiddleware,
	}
}

// Setup user routes
func (s EmployeeRoutes) Setup(handler *gin.RouterGroup) {
	s.logger.Info("Setting up Employee routes")
	api := handler.Group("/employee")
	{
		api.GET("", s.pagination.Handle(), s.userController.ListEmployees)

		api.GET("/type/:type", s.pagination.Handle(), s.userController.ListEmployeesByType)
		api.GET("/deleted", s.userController.ListDeletedEmployee)

		api.GET("/id/:id", s.userController.GetEmployeeByID)
		api.GET("/profile/:id", s.userController.GetProfileByID)

		api.GET("/email/:email", s.userController.GetEmployeeByEmail)
		api.GET("/contact/:contact", s.userController.GetEmployeeByContactNumber)

		api.PATCH("/:id", s.userController.PatchEmployee)
		api.DELETE("/:id", s.userController.DeleteEmployeeByID)
	}
}
