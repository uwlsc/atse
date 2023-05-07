package services

import (
	"magazine_api/component"
	"magazine_api/lib"
	"magazine_api/models/magazine"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// StoryService service layer
type StoryService struct {
	logger lib.Logger
	repo   component.IStoryMgmtComp
}

// NewStoryService creates new instance of StoryService
func NewStoryService(logger lib.Logger, repo component.IStoryMgmtComp) StoryService {
	return StoryService{logger: logger, repo: repo}
}

// Creates the Story in database
func (u StoryService) CreateStory(story *magazine.Story) (*magazine.Story, error) {
	story = u.BeforeCreate(story)

	err := u.repo.CreateStory(*story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// Lists stories form database
func (p StoryService) ListStories(c *gin.Context) ([]*magazine.Story, error) {
	stories, err := p.repo.ListStories()
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// Gets Story by id from database
func (u StoryService) GetStoryById(id uuid.UUID) (*magazine.Story, error) {
	Story, err := u.repo.GetStoryFromID(id)
	if err != nil {
		return nil, err
	}

	return Story, nil
}

// Lists Storys by profile id
func (u StoryService) ListStoriesByProfileId(id uuid.UUID) ([]*magazine.Story, error) {
	Story, err := u.repo.GetStoryFromCreatorId(id)
	if err != nil {
		return nil, err
	}

	return Story, nil
}

// Lists Stories by type
func (u StoryService) ListStoriesByType(story_type string) ([]*magazine.Story, error) {
	Story, err := u.repo.GetStoryFromType(story_type)
	if err != nil {
		return nil, err
	}

	return Story, nil
}

// Update Story by id in our database
func (u StoryService) UpdateStory(id uuid.UUID, patch *map[string]interface{}) error {
	err := u.repo.PatchStory(id, patch)
	if err != nil {
		return err
	}

	return nil
}

// Delete Story by in our database
func (u StoryService) DeleteStory(id uuid.UUID) error {
	err := u.repo.DeleteStory(id)
	if err != nil {
		return err
	}

	return nil
}

// Permanent Delete Story by in our database permanently
func (u StoryService) PermanentDeleteStory(id uuid.UUID) error {
	err := u.repo.PermanentDeleteStory(id)
	if err != nil {
		return err
	}

	return nil
}

func (u StoryService) BeforeCreate(Story *magazine.Story) *magazine.Story {
	Story.ID = uuid.New()
	create := time.Now()
	Story.CreatedOn = &create
	Story.UpdatedOn = &create

	return Story
}
