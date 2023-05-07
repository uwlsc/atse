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

type MagazineIssueHandler struct {
	logger  lib.Logger
	service services.MagazineIssueService
}

func NewMagazineIssueHandler(logger lib.Logger, service services.MagazineIssueService) MagazineIssueHandler {
	return MagazineIssueHandler{
		logger:  logger,
		service: service,
	}
}

// CreateMagazineIssue godoc
// @Summary      Create MagazineIssue
// @Description  It creates MagazineIssue structure
// @Tags         MagazineIssue
// @Accept       json
// @Produce      json
// @Param        MagazineIssue  body      magazine.MagazineIssue  true  "Add MagazineIssue"
// @Success      200       {object}  object{data=magazine.MagazineIssue}
// @Router       /isssue [post]
//
// Creates MagazineIssue
func (s MagazineIssueHandler) CreateMagazineIssue(c *gin.Context) {
	var isssue *magazine.MagazineIssue
	if err := c.ShouldBind(&isssue); err != nil {
		handleError(s.logger, c, err)
		return
	}

	// Create MagazineIssue in our Database
	isssue, err := s.service.CreateMagazineIssue(isssue)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": isssue})
}

// ListAllIssuess godoc
// @Summary      List all stories.
// @Description  List stories
// @Tags         MagazineIssue
// @Produce      json
// @Success      200  {object}  object{data=[]magazine.MagazineIssue}
// @Router       /isssue [get]
//
// List all suppliers from database
func (s MagazineIssueHandler) ListIssues(c *gin.Context) {
	stories, err := s.service.ListIssues(c)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, stories)
}

// ListMagazineIssueFromUserId godoc
// @Summary      Lists MagazineIssue from User Id
// @Description  List MagazineIssue by creator id
// @Tags         MagazineIssue
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=[]magazine.MagazineIssue}
// @Router       /isssue/profile/{id} [get]
//
// List MagazineIssue from creator ID controller
func (a MagazineIssueHandler) ListMagazineIssueByProfileId(c *gin.Context) {

	id := c.Param("id")

	MagazineIssues, err := a.service.ListIssuesByProfileId(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": MagazineIssues})
}

// ListMagazineIssueByType godoc
// @Summary      Lists MagazineIssue Type
// @Description  List MagazineIssue by type
// @Tags         MagazineIssue
// @Produce      json
// @Param        isssue_type   path      string  true  "TYPE"
// @Success      200  {object}  object{data=[]magazine.MagazineIssue}
// @Router       /isssue/type/{isssue_type} [get]
//
// List MagazineIssue from creator ID controller
func (a MagazineIssueHandler) ListIssuesByType(c *gin.Context) {

	isssue_type := c.Param("type")

	MagazineIssues, err := a.service.ListIssuesByType(isssue_type)

	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": MagazineIssues})
}

// GetMagazineIssueById godoc
// @Summary      Gets One MagazineIssue by ID
// @Description  Gets One MagazineIssue by D
// @Tags         MagazineIssue
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=models.MagazineIssue}
// @Router       /isssue/id/{id} [get]
//
// Gets MagazineIssue By Company ID controller
func (a MagazineIssueHandler) GetMagazineIssueById(c *gin.Context) {
	id := c.Param("id")

	MagazineIssue, err := a.service.GetMagazineIssueById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": MagazineIssue})
}

// UpdateMagazineIssue godoc
// @Summary      Update MagazineIssue
// @Description  Updates MagazineIssue of employee
// @Tags         MagazineIssue
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Update MagazineIssue"
// @Success      200  {object}  object{data=models.MagazineIssue}
// @Router       /isssue/{id} [patch]
// Patch MagazineIssue of creator by Id controller
func (a MagazineIssueHandler) PatchMagazineIssueById(c *gin.Context) {
	id := c.Param("id")

	MagazineIssue, err := a.service.GetMagazineIssueById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	var newMagazineIssue magazine.IssuseBase
	if err := c.ShouldBindJSON(&newMagazineIssue); err != nil {
		handleError(a.logger, c, err)
		return
	}

	MagazineIssueMap := structomap.New().UseSnakeCase().PickAll().
		Omit("CreatorId").
		OmitIf(func(ch interface{}) bool {
			return newMagazineIssue.IssueCode == nil
		}, "IssueCode").
		OmitIf(func(ch interface{}) bool {
			return newMagazineIssue.ContentCode == nil
		}, "ContentCode").
		OmitIf(func(ch interface{}) bool {
			return newMagazineIssue.AdvertCode == nil
		}, "AdvertCode").
		OmitIf(func(ch interface{}) bool {
			return newMagazineIssue.Remarks == nil
		}, "Remarks").
		Transform(newMagazineIssue)

	if len(MagazineIssueMap) > 0 {
		MagazineIssueMap["updated_on"] = time.Now()
		MagazineIssueMap["id"] = MagazineIssue.ID

		err := a.service.UpdateMagazineIssue(MagazineIssue.ID, &MagazineIssueMap)
		if err != nil {
			handleError(a.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": MagazineIssueMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}

// DeleteMagazineIssue godoc
// @Summary      Soft Delete an MagazineIssue
// @Description  Delete by MagazineIssueID
// @Tags         MagazineIssue
// @Produce      json
// @Param        id   path      string  true  "Unique ID"
// @Success      204  {object}  object{data=string}
// @Router       /isssue/{id} [delete]
//
// Delete MagazineIssue By ID controller
func (a MagazineIssueHandler) DeleteMagazineIssueByID(c *gin.Context) {
	id := c.Param("id")

	err := a.service.DeleteMagazineIssue(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}
