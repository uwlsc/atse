package handlers

import (
	"magazine_api/lib"
	"magazine_api/models/magazine"
	"magazine_api/services"
	"time"

	"github.com/danhper/structomap"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PhotoHandler struct {
	logger  lib.Logger
	service services.PhotoService
}

func NewPhotoHandler(logger lib.Logger, service services.PhotoService) PhotoHandler {
	return PhotoHandler{
		logger:  logger,
		service: service,
	}
}

// CreatePhoto godoc
// @Summary      Create Photo
// @Description  It creates Photo structure
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Param        Photo  body      magazine.Photograph  true  "Add Photo"
// @Success      200       {object}  object{data=magazine.Photograph}
// @Router       /photo [post]
//
// Creates Photo
func (s PhotoHandler) CreatePhoto(c *gin.Context) {
	var photo *magazine.Photograph
	if err := c.ShouldBind(&photo); err != nil {
		handleError(s.logger, c, err)
		return
	}

	// Create Photo in our Database
	photo, err := s.service.CreatePhoto(photo)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": photo})
}

// ListAllPhotos godoc
// @Summary      List all stories.
// @Description  List stories
// @Tags         Photo
// @Produce      json
// @Success      200  {object}  object{data=[]magazine.Photograph}
// @Router       /photo [get]
//
// List all suppliers from database
func (s PhotoHandler) ListPhotos(c *gin.Context) {
	stories, err := s.service.ListPhotos(c)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, stories)
}

// ListPhotoFromUserId godoc
// @Summary      Lists Photo from User Id
// @Description  List Photo by creator id
// @Tags         Photo
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=[]magazine.Photograph}
// @Router       /photo/profile/{id} [get]
//
// List Photo from creator ID controller
func (a PhotoHandler) ListPhotoByProfileId(c *gin.Context) {

	id := c.Param("id")

	Photos, err := a.service.ListPhotosByProfileId(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Photos})
}

// ListPhotoByType godoc
// @Summary      Lists Photo Type
// @Description  List Photo by type
// @Tags         Photo
// @Produce      json
// @Param        photo_type   path      string  true  "TYPE"
// @Success      200  {object}  object{data=[]magazine.Photograph}
// @Router       /photo/type/{photo_type} [get]
//
// List Photo from creator ID controller
func (a PhotoHandler) ListPhotosByType(c *gin.Context) {

	photo_type := c.Param("type")

	Photos, err := a.service.ListPhotosByType(photo_type)

	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Photos})
}

// GetPhotoById godoc
// @Summary      Gets One Photo by ID
// @Description  Gets One Photo by D
// @Tags         Photo
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=models.Photograph}
// @Router       /photo/id/{id} [get]
//
// Gets Photo By Company ID controller
func (a PhotoHandler) GetPhotoById(c *gin.Context) {
	id := c.Param("id")

	Photo, err := a.service.GetPhotoById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Photo})
}

// UpdatePhoto godoc
// @Summary      Update Photo
// @Description  Updates Photo of employee
// @Tags         Photo
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Update Photo"
// @Success      200  {object}  object{data=models.Photograph}
// @Router       /photo/{id} [patch]
// Patch Photo of creator by Id controller
func (a PhotoHandler) PatchPhotoById(c *gin.Context) {
	id := c.Param("id")

	Photo, err := a.service.GetPhotoById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	var newPhoto magazine.PhotographBase
	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		handleError(a.logger, c, err)
		return
	}

	PhotoMap := structomap.New().UseSnakeCase().PickAll().
		Omit("CreatorId").
		OmitIf(func(ch interface{}) bool {
			return newPhoto.PhotographID == nil
		}, "PhotographID").
		OmitIf(func(ch interface{}) bool {
			return newPhoto.PhotographCode == nil
		}, "PhotographCode").
		OmitIf(func(ch interface{}) bool {
			return newPhoto.PhotographTitle == nil
		}, "PhotographTitle").
		OmitIf(func(ch interface{}) bool {
			return newPhoto.PhotographType == nil
		}, "PhotographType").
		OmitIf(func(ch interface{}) bool {
			return newPhoto.DocumentURL == nil
		}, "DocumentURL").
		OmitIf(func(ch interface{}) bool {
			return newPhoto.Remarks == nil
		}, "Remarks").
		Transform(newPhoto)

	if len(PhotoMap) > 0 {
		PhotoMap["updated_on"] = time.Now()
		PhotoMap["id"] = Photo.ID

		err := a.service.UpdatePhoto(Photo.ID, &PhotoMap)
		if err != nil {
			handleError(a.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": PhotoMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}

// DeletePhoto godoc
// @Summary      Soft Delete an Photo
// @Description  Delete by PhotoID
// @Tags         Photo
// @Produce      json
// @Param        id   path      string  true  "Unique ID"
// @Success      204  {object}  object{data=string}
// @Router       /photo/{id} [delete]
//
// Delete Photo By ID controller
func (a PhotoHandler) DeletePhotoByID(c *gin.Context) {
	id := c.Param("id")

	err := a.service.DeletePhoto(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}
