package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

type AdvertRoutes struct {
	logger        lib.Logger
	handler       infrastructure.Router
	pagination    middlewares.PaginationMiddleware
	advertHandler handlers.AdvertHandler
}

func NewAdvertRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	advertHandler handlers.AdvertHandler) AdvertRoutes {
	return AdvertRoutes{
		handler:       handler,
		logger:        logger,
		pagination:    pagination,
		advertHandler: advertHandler,
	}
}

// Setup salary routes
func (a AdvertRoutes) Setup(handler *gin.RouterGroup) {
	a.logger.Info("Setting up Document routes")
	api := handler.Group("/advert")
	{
		api.POST("", a.advertHandler.CreateAdvert)
		api.GET("", a.pagination.Handle(), a.advertHandler.ListAdvert)
		api.GET("/profile/:id", a.advertHandler.ListAdvertByProfileId)

		api.PATCH("/:id", a.advertHandler.PatchAdvertById)
		api.DELETE("/:id", a.advertHandler.DeleteAdvertByID)
	}
}
