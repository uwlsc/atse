package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

type MagazineRoutes struct {
	logger          lib.Logger
	handler         infrastructure.Router
	pagination      middlewares.PaginationMiddleware
	magazineHandler handlers.MagazineHandler
}

func NewMagazineRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	magazineHandler handlers.MagazineHandler) MagazineRoutes {
	return MagazineRoutes{
		handler:         handler,
		logger:          logger,
		pagination:      pagination,
		magazineHandler: magazineHandler,
	}
}

// Setup salary routes
func (a MagazineRoutes) Setup(handler *gin.RouterGroup) {
	a.logger.Info("Setting up Magazine routes")
	api := handler.Group("/magazine")
	{
		api.POST("", a.magazineHandler.CreateMagazine)
		api.GET("", a.pagination.Handle(), a.magazineHandler.ListMagazines)
		api.GET("/profile/:id", a.magazineHandler.ListMagazineByProfileId)
		api.GET("/type/:magazine_type", a.magazineHandler.ListMagazinesByType)

		api.PATCH("/:id", a.magazineHandler.PatchMagazineById)
		api.DELETE("/:id", a.magazineHandler.DeleteMagazineByID)
	}
}
