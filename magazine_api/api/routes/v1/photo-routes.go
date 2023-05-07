package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

type PhotoRoutes struct {
	logger       lib.Logger
	handler      infrastructure.Router
	pagination   middlewares.PaginationMiddleware
	photoHandler handlers.PhotoHandler
}

func NewPhotoRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	photoHandler handlers.PhotoHandler) PhotoRoutes {
	return PhotoRoutes{
		handler:      handler,
		logger:       logger,
		pagination:   pagination,
		photoHandler: photoHandler,
	}
}

// Setup salary routes
func (a PhotoRoutes) Setup(handler *gin.RouterGroup) {
	a.logger.Info("Setting up Photo routes")
	api := handler.Group("/photo")
	{
		api.POST("", a.photoHandler.CreatePhoto)
		api.GET("", a.pagination.Handle(), a.photoHandler.ListPhotos)
		api.GET("/profile/:id", a.photoHandler.ListPhotoByProfileId)
		api.GET("/type/:photo_type", a.photoHandler.ListPhotosByType)

		api.PATCH("/:id", a.photoHandler.PatchPhotoById)
		api.DELETE("/:id", a.photoHandler.DeletePhotoByID)
	}
}
