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

type MagazineHandler struct {
	logger  lib.Logger
	service services.MagazineService
}

func NewMagazineHandler(logger lib.Logger, service services.MagazineService) MagazineHandler {
	return MagazineHandler{
		logger:  logger,
		service: service,
	}
}

// CreateMagazine godoc
// @Summary      Create Magazine
// @Description  It creates Magazine structure
// @Tags         Magazine
// @Accept       json
// @Produce      json
// @Param        Magazine  body      magazine.Magazine  true  "Add Magazine"
// @Success      200       {object}  object{data=magazine.Magazine}
// @Router       /magazine [post]
//
// Creates Magazine
func (s MagazineHandler) CreateMagazine(c *gin.Context) {
	var magazine *magazine.Magazine
	if err := c.ShouldBind(&magazine); err != nil {
		handleError(s.logger, c, err)
		return
	}

	// Create Magazine in our Database
	magazine, err := s.service.CreateMagazine(magazine)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": magazine})
}

// ListAllMagaziness godoc
// @Summary      List all stories.
// @Description  List stories
// @Tags         Magazine
// @Produce      json
// @Success      200  {object}  object{data=[]magazine.Magazine}
// @Router       /magazine [get]
//
// List all suppliers from database
func (s MagazineHandler) ListMagazines(c *gin.Context) {
	stories, err := s.service.ListMagazines(c)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, stories)
}

// ListMagazineFromUserId godoc
// @Summary      Lists Magazine from User Id
// @Description  List Magazine by creator id
// @Tags         Magazine
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=[]magazine.Magazine}
// @Router       /magazine/profile/{id} [get]
//
// List Magazine from creator ID controller
func (a MagazineHandler) ListMagazineByProfileId(c *gin.Context) {

	id := c.Param("id")

	Magazines, err := a.service.ListMagazinesByProfileId(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Magazines})
}

// ListMagazineByType godoc
// @Summary      Lists Magazine Type
// @Description  List Magazine by type
// @Tags         Magazine
// @Produce      json
// @Param        magazine_type   path      string  true  "TYPE"
// @Success      200  {object}  object{data=[]magazine.Magazine}
// @Router       /magazine/type/{magazine_type} [get]
//
// List Magazine from creator ID controller
func (a MagazineHandler) ListMagazinesByType(c *gin.Context) {

	magazine_type := c.Param("type")

	Magazines, err := a.service.ListMagazinesByType(magazine_type)

	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Magazines})
}

// GetMagazineById godoc
// @Summary      Gets One Magazine by ID
// @Description  Gets One Magazine by D
// @Tags         Magazine
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=models.Magazine}
// @Router       /magazine/id/{id} [get]
//
// Gets Magazine By Company ID controller
func (a MagazineHandler) GetMagazineById(c *gin.Context) {
	id := c.Param("id")

	Magazine, err := a.service.GetMagazineById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Magazine})
}

// UpdateMagazine godoc
// @Summary      Update Magazine
// @Description  Updates Magazine of employee
// @Tags         Magazine
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Update Magazine"
// @Success      200  {object}  object{data=models.Magazine}
// @Router       /magazine/{id} [patch]
// Patch Magazine of creator by Id controller
func (a MagazineHandler) PatchMagazineById(c *gin.Context) {
	id := c.Param("id")

	Magazine, err := a.service.GetMagazineById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	var newMagazine magazine.MagazineBase
	if err := c.ShouldBindJSON(&newMagazine); err != nil {
		handleError(a.logger, c, err)
		return
	}

	MagazineMap := structomap.New().UseSnakeCase().PickAll().
		Omit("CreatorId").
		OmitIf(func(ch interface{}) bool {
			return newMagazine.MagazineCode == nil
		}, "MagazineCode").
		OmitIf(func(ch interface{}) bool {
			return newMagazine.IssueCode == nil
		}, "IssueCode").
		OmitIf(func(ch interface{}) bool {
			return newMagazine.Placement == nil
		}, "Placement").
		OmitIf(func(ch interface{}) bool {
			return newMagazine.Remarks == nil
		}, "Remarks").
		Transform(newMagazine)

	if len(MagazineMap) > 0 {
		MagazineMap["updated_on"] = time.Now()
		MagazineMap["id"] = Magazine.ID

		err := a.service.UpdateMagazine(Magazine.ID, &MagazineMap)
		if err != nil {
			handleError(a.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": MagazineMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}

// DeleteMagazine godoc
// @Summary      Soft Delete an Magazine
// @Description  Delete by MagazineID
// @Tags         Magazine
// @Produce      json
// @Param        id   path      string  true  "Unique ID"
// @Success      204  {object}  object{data=string}
// @Router       /magazine/{id} [delete]
//
// Delete Magazine By ID controller
func (a MagazineHandler) DeleteMagazineByID(c *gin.Context) {
	id := c.Param("id")

	err := a.service.DeleteMagazine(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}
