package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

type MagazineIssueRoutes struct {
	logger       lib.Logger
	handler      infrastructure.Router
	pagination   middlewares.PaginationMiddleware
	issueHandler handlers.MagazineIssueHandler
}

func NewMagazineIssueRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	issueHandler handlers.MagazineIssueHandler) MagazineIssueRoutes {
	return MagazineIssueRoutes{
		handler:      handler,
		logger:       logger,
		pagination:   pagination,
		issueHandler: issueHandler,
	}
}

// Setup salary routes
func (a MagazineIssueRoutes) Setup(handler *gin.RouterGroup) {
	a.logger.Info("Setting up MagazineIssue routes")
	api := handler.Group("/issue")
	{
		api.POST("", a.issueHandler.CreateMagazineIssue)
		api.GET("", a.pagination.Handle(), a.issueHandler.ListIssues)
		api.GET("/profile/:id", a.issueHandler.ListMagazineIssueByProfileId)
		api.GET("/type/:issue_type", a.issueHandler.ListIssuesByType)

		api.PATCH("/:id", a.issueHandler.PatchMagazineIssueById)
		api.DELETE("/:id", a.issueHandler.DeleteMagazineIssueByID)
	}
}
