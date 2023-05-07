package handlers

import (
	"magazine_api/lib"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewUserHandler),
	fx.Provide(NewEmployeeHandler),
	fx.Provide(NewUploadHandler),
	fx.Provide(NewTransactionHandler),
	fx.Provide(NewUserProfileHandler),
	fx.Provide(NewStoryHandler),
	fx.Provide(NewAdvertHandler),
	fx.Provide(NewContentHandler),
	fx.Provide(NewMagazineIssueHandler),
	fx.Provide(NewMagazineHandler),
	fx.Provide(NewPhotoHandler),
)

func handleError(logger lib.Logger, c *gin.Context, err error) {
	logger.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
