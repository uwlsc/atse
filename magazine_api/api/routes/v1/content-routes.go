package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

type ContentRoutes struct {
	logger         lib.Logger
	handler        infrastructure.Router
	pagination     middlewares.PaginationMiddleware
	contentHandler handlers.ContentHandler
}

func NewContentRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	contentHandler handlers.ContentHandler) ContentRoutes {
	return ContentRoutes{
		handler:        handler,
		logger:         logger,
		pagination:     pagination,
		contentHandler: contentHandler,
	}
}

// Setup salary routes
func (a ContentRoutes) Setup(handler *gin.RouterGroup) {
	a.logger.Info("Setting up Content routes")
	api := handler.Group("/content")
	{
		api.POST("", a.contentHandler.CreateContent)
		api.GET("", a.pagination.Handle(), a.contentHandler.ListContents)
		api.GET("/profile/:id", a.contentHandler.ListContentByProfileId)
		api.GET("/type/:content_type", a.contentHandler.ListContentsByType)

		api.PATCH("/:id", a.contentHandler.PatchContentById)
		api.DELETE("/:id", a.contentHandler.DeleteContentByID)
	}
}
