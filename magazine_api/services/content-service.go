package services

import (
	"magazine_api/component"
	"magazine_api/lib"
	"magazine_api/models/magazine"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ContentService service layer
type ContentService struct {
	logger lib.Logger
	repo   component.IContentMgmtComp
}

// NewContentService creates new instance of ContentService
func NewContentService(logger lib.Logger, repo component.IContentMgmtComp) ContentService {
	return ContentService{logger: logger, repo: repo}
}

// Creates the Content in database
func (u ContentService) CreateContent(story *magazine.Content) (*magazine.Content, error) {
	story = u.BeforeCreate(story)

	err := u.repo.CreateContent(*story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// Lists stories form database
func (p ContentService) ListContents(c *gin.Context) ([]*magazine.Content, error) {
	stories, err := p.repo.ListContents()
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// Gets Content by id from database
func (u ContentService) GetContentById(id uuid.UUID) (*magazine.Content, error) {
	Content, err := u.repo.GetContentFromID(id)
	if err != nil {
		return nil, err
	}

	return Content, nil
}

// Lists Contents by profile id
func (u ContentService) ListContentsByProfileId(id uuid.UUID) ([]*magazine.Content, error) {
	Content, err := u.repo.GetContentFromCreatorId(id)
	if err != nil {
		return nil, err
	}

	return Content, nil
}

// Lists Contents by type
func (u ContentService) ListContentsByType(story_type string) ([]*magazine.Content, error) {
	Content, err := u.repo.GetContentFromType(story_type)
	if err != nil {
		return nil, err
	}

	return Content, nil
}

// Update Content by id in our database
func (u ContentService) UpdateContent(id uuid.UUID, patch *map[string]interface{}) error {
	err := u.repo.PatchContent(id, patch)
	if err != nil {
		return err
	}

	return nil
}

// Delete Content by in our database
func (u ContentService) DeleteContent(id uuid.UUID) error {
	err := u.repo.DeleteContent(id)
	if err != nil {
		return err
	}

	return nil
}

// Permanent Delete Content by in our database permanently
func (u ContentService) PermanentDeleteContent(id uuid.UUID) error {
	err := u.repo.PermanentDeleteContent(id)
	if err != nil {
		return err
	}

	return nil
}

func (u ContentService) BeforeCreate(Content *magazine.Content) *magazine.Content {
	Content.ID = uuid.New()
	create := time.Now()
	Content.CreatedOn = &create
	Content.UpdatedOn = &create

	return Content
}
