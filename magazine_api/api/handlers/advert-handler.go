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

type AdvertHandler struct {
	logger  lib.Logger
	service services.AdvertService
}

func NewAdvertHandler(logger lib.Logger, service services.AdvertService) AdvertHandler {
	return AdvertHandler{
		logger:  logger,
		service: service,
	}
}

// CreateAdvert godoc
// @Summary      Create Advert
// @Description  It creates Advert structure
// @Tags         Advert
// @Accept       json
// @Produce      json
// @Param        Advert  body      magazine.Advert  true  "Add Advert"
// @Success      200       {object}  object{data=magazine.Advert}
// @Router       /ad [post]
//
// Creates Advert
func (s AdvertHandler) CreateAdvert(c *gin.Context) {
	var ad *magazine.Advert
	if err := c.ShouldBind(&ad); err != nil {
		handleError(s.logger, c, err)
		return
	}

	// Create Advert in our Database
	ad, err := s.service.CreateAdvert(ad)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": ad})
}

// ListAllAdverts godoc
// @Summary      List all stories.
// @Description  List stories
// @Tags         Advert
// @Produce      json
// @Success      200  {object}  object{data=[]magazine.Advert}
// @Router       /ad [get]
//
// List all suppliers from database
func (s AdvertHandler) ListAdvert(c *gin.Context) {
	stories, err := s.service.ListAds(c)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, stories)
}

// ListAdvertFromUserId godoc
// @Summary      Lists Advert from User Id
// @Description  List Advert by creator id
// @Tags         Advert
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=[]magazine.Advert}
// @Router       /ad/profile/{id} [get]
//
// List Advert from creator ID controller
func (a AdvertHandler) ListAdvertByProfileId(c *gin.Context) {

	id := c.Param("id")

	Advert, err := a.service.ListAdsByProfileId(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Advert})
}

// GetAdvertById godoc
// @Summary      Gets One Advert by ID
// @Description  Gets One Advert by D
// @Tags         Advert
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=models.Advert}
// @Router       /ad/id/{id} [get]
//
// Gets Advert By Company ID controller
func (a AdvertHandler) GetAdvertById(c *gin.Context) {
	id := c.Param("id")

	Advert, err := a.service.GetAdvertById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Advert})
}

// UpdateAdvert godoc
// @Summary      Update Advert
// @Description  Updates Advert of employee
// @Tags         Advert
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Update Advert"
// @Success      200  {object}  object{data=models.Advert}
// @Router       /ad/{id} [patch]
// Patch Advert of creator by Id controller
func (a AdvertHandler) PatchAdvertById(c *gin.Context) {
	id := c.Param("id")

	Advert, err := a.service.GetAdvertById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	var newAdvert magazine.AdvertBase
	if err := c.ShouldBindJSON(&newAdvert); err != nil {
		handleError(a.logger, c, err)
		return
	}

	AdvertMap := structomap.New().UseSnakeCase().PickAll().
		Omit("CreatorId").
		OmitIf(func(ch interface{}) bool {
			return newAdvert.AdvertCode == nil
		}, "AdvertCode").
		OmitIf(func(ch interface{}) bool {
			return newAdvert.AdvertTitle == nil
		}, "AdvertTitle").
		OmitIf(func(ch interface{}) bool {
			return newAdvert.AdvertType == nil
		}, "AdvertType").
		OmitIf(func(ch interface{}) bool {
			return newAdvert.AdvertContent == nil
		}, "AdvertContent").
		Transform(newAdvert)

	if len(AdvertMap) > 0 {
		AdvertMap["updated_on"] = time.Now()
		AdvertMap["id"] = Advert.ID

		err := a.service.UpdateAdvert(Advert.ID, &AdvertMap)
		if err != nil {
			handleError(a.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": AdvertMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}

// DeleteAdvert godoc
// @Summary      Soft Delete an Advert
// @Description  Delete by AdvertID
// @Tags         Advert
// @Produce      json
// @Param        id   path      string  true  "Unique ID"
// @Success      204  {object}  object{data=string}
// @Router       /ad/{id} [delete]
//
// Delete Advert By ID controller
func (a AdvertHandler) DeleteAdvertByID(c *gin.Context) {
	id := c.Param("id")

	err := a.service.DeleteAdvert(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}
