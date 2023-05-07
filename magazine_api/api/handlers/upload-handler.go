package handlers

import (
	"errors"
	"magazine_api/constants"
	"magazine_api/lib"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	logger lib.Logger
}

func NewUploadHandler(logger lib.Logger) UploadHandler {
	return UploadHandler{logger: logger}
}

// UploadSingleFile godoc
// @Summary      Upload Single File
// @Description  It uploads the file to S3
// @Tags         Upload
// @Accept       json
// @Produce      json
// @Param        file  formData  file  true  "Upload file"
// @Success      200
// @Router       /upload [post]
//
// Upload files
func (u UploadHandler) UploadFile(c *gin.Context) {

	metadata, isFile := c.Get(constants.File)

	if isFile {
		file_metadata := metadata.(lib.UploadedFiles)
		url := file_metadata[0].URL

		c.JSON(200, gin.H{"url": url})
		return
	}

	handleError(u.logger, c, errors.New("no image uploaded"))
}
