package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

type StoryRoutes struct {
	logger       lib.Logger
	handler      infrastructure.Router
	pagination   middlewares.PaginationMiddleware
	storyHandler handlers.StoryHandler
}

func NewStoryRoutes(logger lib.Logger,
	handler infrastructure.Router,
	pagination middlewares.PaginationMiddleware,
	storyHandler handlers.StoryHandler) StoryRoutes {
	return StoryRoutes{
		handler:      handler,
		logger:       logger,
		pagination:   pagination,
		storyHandler: storyHandler,
	}
}

// Setup salary routes
func (a StoryRoutes) Setup(handler *gin.RouterGroup) {
	a.logger.Info("Setting up Story routes")
	api := handler.Group("/story")
	{
		api.POST("", a.storyHandler.CreateStory)
		api.GET("", a.pagination.Handle(), a.storyHandler.ListStories)
		api.GET("/profile/:id", a.storyHandler.ListStoryByProfileId)
		api.GET("/type/:story_type", a.storyHandler.ListStoriesByType)

		api.PATCH("/:id", a.storyHandler.PatchStoryById)
		api.DELETE("/:id", a.storyHandler.DeleteStoryByID)
	}
}
