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

type ContentHandler struct {
	logger  lib.Logger
	service services.ContentService
}

func NewContentHandler(logger lib.Logger, service services.ContentService) ContentHandler {
	return ContentHandler{
		logger:  logger,
		service: service,
	}
}

// CreateContent godoc
// @Summary      Create Content
// @Description  It creates Content structure
// @Tags         Content
// @Accept       json
// @Produce      json
// @Param        Content  body      magazine.Content  true  "Add Content"
// @Success      200       {object}  object{data=magazine.Content}
// @Router       /content [post]
//
// Creates Content
func (s ContentHandler) CreateContent(c *gin.Context) {
	var content *magazine.Content
	if err := c.ShouldBind(&content); err != nil {
		handleError(s.logger, c, err)
		return
	}

	// Create Content in our Database
	content, err := s.service.CreateContent(content)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": content})
}

// ListAllContentss godoc
// @Summary      List all contents.
// @Description  List contents
// @Tags         Content
// @Produce      json
// @Success      200  {object}  object{data=[]magazine.Content}
// @Router       /content [get]
//
// List all suppliers from database
func (s ContentHandler) ListContents(c *gin.Context) {
	contents, err := s.service.ListContents(c)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, contents)
}

// ListContentFromUserId godoc
// @Summary      Lists Content from User Id
// @Description  List Content by creator id
// @Tags         Content
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=[]magazine.Content}
// @Router       /content/profile/{id} [get]
//
// List Content from creator ID controller
func (a ContentHandler) ListContentByProfileId(c *gin.Context) {

	id := c.Param("id")

	Contents, err := a.service.ListContentsByProfileId(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Contents})
}

// ListContentByType godoc
// @Summary      Lists Content Type
// @Description  List Content by type
// @Tags         Content
// @Produce      json
// @Param        content_type   path      string  true  "TYPE"
// @Success      200  {object}  object{data=[]magazine.Content}
// @Router       /content/type/{content_type} [get]
//
// List Content from creator ID controller
func (a ContentHandler) ListContentsByType(c *gin.Context) {

	content_type := c.Param("type")

	Contents, err := a.service.ListContentsByType(content_type)

	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Contents})
}

// GetContentById godoc
// @Summary      Gets One Content by ID
// @Description  Gets One Content by D
// @Tags         Content
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=models.Content}
// @Router       /content/id/{id} [get]
//
// Gets Content By Company ID controller
func (a ContentHandler) GetContentById(c *gin.Context) {
	id := c.Param("id")

	Content, err := a.service.GetContentById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Content})
}

// UpdateContent godoc
// @Summary      Update Content
// @Description  Updates Content of employee
// @Tags         Content
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Update Content"
// @Success      200  {object}  object{data=models.Content}
// @Router       /content/{id} [patch]
// Patch Content of creator by Id controller
func (a ContentHandler) PatchContentById(c *gin.Context) {
	id := c.Param("id")

	Content, err := a.service.GetContentById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	var newContent magazine.ContentBase
	if err := c.ShouldBindJSON(&newContent); err != nil {
		handleError(a.logger, c, err)
		return
	}

	ContentMap := structomap.New().UseSnakeCase().PickAll().
		Omit("CreatorId").
		OmitIf(func(ch interface{}) bool {
			return newContent.ContentCode == nil
		}, "ContentCode").
		OmitIf(func(ch interface{}) bool {
			return newContent.StoryCode == nil
		}, "StoryCode").
		OmitIf(func(ch interface{}) bool {
			return newContent.PhotographCode == nil
		}, "PhotographCode").
		OmitIf(func(ch interface{}) bool {
			return newContent.Remarks == nil
		}, "Remarks").
		Transform(newContent)

	if len(ContentMap) > 0 {
		ContentMap["updated_on"] = time.Now()
		ContentMap["id"] = Content.ID

		err := a.service.UpdateContent(Content.ID, &ContentMap)
		if err != nil {
			handleError(a.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": ContentMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}

// DeleteContent godoc
// @Summary      Soft Delete an Content
// @Description  Delete by ContentID
// @Tags         Content
// @Produce      json
// @Param        id   path      string  true  "Unique ID"
// @Success      204  {object}  object{data=string}
// @Router       /content/{id} [delete]
//
// Delete Content By ID controller
func (a ContentHandler) DeleteContentByID(c *gin.Context) {
	id := c.Param("id")

	err := a.service.DeleteContent(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}
