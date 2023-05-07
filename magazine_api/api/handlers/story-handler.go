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

type StoryHandler struct {
	logger  lib.Logger
	service services.StoryService
}

func NewStoryHandler(logger lib.Logger, service services.StoryService) StoryHandler {
	return StoryHandler{
		logger:  logger,
		service: service,
	}
}

// CreateStory godoc
// @Summary      Create Story
// @Description  It creates Story structure
// @Tags         Story
// @Accept       json
// @Produce      json
// @Param        Story  body      magazine.Story  true  "Add Story"
// @Success      200       {object}  object{data=magazine.Story}
// @Router       /story [post]
//
// Creates Story
func (s StoryHandler) CreateStory(c *gin.Context) {
	var story *magazine.Story
	if err := c.ShouldBind(&story); err != nil {
		handleError(s.logger, c, err)
		return
	}

	// Create Story in our Database
	story, err := s.service.CreateStory(story)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": story})
}

// ListAllStoriess godoc
// @Summary      List all stories.
// @Description  List stories
// @Tags         Story
// @Produce      json
// @Success      200  {object}  object{data=[]magazine.Story}
// @Router       /story [get]
//
// List all suppliers from database
func (s StoryHandler) ListStories(c *gin.Context) {
	stories, err := s.service.ListStories(c)
	if err != nil {
		handleError(s.logger, c, err)
		return
	}

	c.JSON(200, stories)
}

// ListStoryFromUserId godoc
// @Summary      Lists Story from User Id
// @Description  List Story by creator id
// @Tags         Story
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=[]magazine.Story}
// @Router       /story/profile/{id} [get]
//
// List Story from creator ID controller
func (a StoryHandler) ListStoryByProfileId(c *gin.Context) {

	id := c.Param("id")

	Storys, err := a.service.ListStoriesByProfileId(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Storys})
}

// ListStoryByType godoc
// @Summary      Lists Story Type
// @Description  List Story by type
// @Tags         Story
// @Produce      json
// @Param        story_type   path      string  true  "TYPE"
// @Success      200  {object}  object{data=[]magazine.Story}
// @Router       /story/type/{story_type} [get]
//
// List Story from creator ID controller
func (a StoryHandler) ListStoriesByType(c *gin.Context) {

	story_type := c.Param("type")

	Storys, err := a.service.ListStoriesByType(story_type)

	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Storys})
}

// GetStoryById godoc
// @Summary      Gets One Story by ID
// @Description  Gets One Story by D
// @Tags         Story
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  object{data=models.Story}
// @Router       /story/id/{id} [get]
//
// Gets Story By Company ID controller
func (a StoryHandler) GetStoryById(c *gin.Context) {
	id := c.Param("id")

	Story, err := a.service.GetStoryById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": Story})
}

// UpdateStory godoc
// @Summary      Update Story
// @Description  Updates Story of employee
// @Tags         Story
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Update Story"
// @Success      200  {object}  object{data=models.Story}
// @Router       /story/{id} [patch]
// Patch Story of creator by Id controller
func (a StoryHandler) PatchStoryById(c *gin.Context) {
	id := c.Param("id")

	Story, err := a.service.GetStoryById(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	var newStory magazine.StoryBase
	if err := c.ShouldBindJSON(&newStory); err != nil {
		handleError(a.logger, c, err)
		return
	}

	StoryMap := structomap.New().UseSnakeCase().PickAll().
		Omit("CreatorId").
		OmitIf(func(ch interface{}) bool {
			return newStory.StoryCode == nil
		}, "StoryCode").
		OmitIf(func(ch interface{}) bool {
			return newStory.StoryTitle == nil
		}, "StoryTitle").
		OmitIf(func(ch interface{}) bool {
			return newStory.StoryType == nil
		}, "StoryType").
		OmitIf(func(ch interface{}) bool {
			return newStory.StoryContent == nil
		}, "StoryContent").
		Transform(newStory)

	if len(StoryMap) > 0 {
		StoryMap["updated_on"] = time.Now()
		StoryMap["id"] = Story.ID

		err := a.service.UpdateStory(Story.ID, &StoryMap)
		if err != nil {
			handleError(a.logger, c, err)
			return
		}

		c.JSON(200, gin.H{"data": StoryMap})
		return
	}

	c.JSON(200, gin.H{"data": "nothing to update"})
}

// DeleteStory godoc
// @Summary      Soft Delete an Story
// @Description  Delete by StoryID
// @Tags         Story
// @Produce      json
// @Param        id   path      string  true  "Unique ID"
// @Success      204  {object}  object{data=string}
// @Router       /story/{id} [delete]
//
// Delete Story By ID controller
func (a StoryHandler) DeleteStoryByID(c *gin.Context) {
	id := c.Param("id")

	err := a.service.DeleteStory(uuid.MustParse(id))
	if err != nil {
		handleError(a.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "successfully deleted"})
}
