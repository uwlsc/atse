package v1

import (
	"magazine_api/api/handlers"
	"magazine_api/api/middlewares"
	"magazine_api/infrastructure"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

// UploadRoutes struct
type UploadRoutes struct {
	logger           lib.Logger
	handler          infrastructure.Router
	uploadMiddleware middlewares.UploadMiddleware
	uploadHandler    handlers.UploadHandler
}

func NewUploadRoutes(logger lib.Logger,
	handler infrastructure.Router,
	uploadMiddleware middlewares.UploadMiddleware,
	uploadHandler handlers.UploadHandler) UploadRoutes {
	return UploadRoutes{
		handler:          handler,
		logger:           logger,
		uploadMiddleware: uploadMiddleware,
		uploadHandler:    uploadHandler,
	}
}

// Setup user routes
func (s UploadRoutes) Setup(handler *gin.RouterGroup) {
	s.logger.Info("Setting up Upload routes")
	api := handler.Group("/upload")
	{
		api.POST("", s.uploadMiddleware.Push(
			s.uploadMiddleware.Config().
				WebpEnable(true).
				ThumbEnable(true).
				Folder("docs_upload")).
			Handle(), s.uploadHandler.UploadFile)
	}
}
