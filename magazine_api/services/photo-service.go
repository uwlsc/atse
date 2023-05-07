package services

import (
	"magazine_api/component"
	"magazine_api/lib"
	"magazine_api/models/magazine"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PhotoService service layer
type PhotoService struct {
	logger lib.Logger
	repo   component.IPhotographMgmtComp
}

// NewPhotoService creates new instance of PhotoService
func NewPhotoService(logger lib.Logger, repo component.IPhotographMgmtComp) PhotoService {
	return PhotoService{logger: logger, repo: repo}
}

// Creates the Photo in database
func (u PhotoService) CreatePhoto(story *magazine.Photograph) (*magazine.Photograph, error) {
	story = u.BeforeCreate(story)

	err := u.repo.CreatePhoto(*story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// Lists photos form database
func (p PhotoService) ListPhotos(c *gin.Context) ([]*magazine.Photograph, error) {
	photos, err := p.repo.ListPhotos()
	if err != nil {
		return nil, err
	}
	return photos, nil
}

// Gets Photo by id from database
func (u PhotoService) GetPhotoById(id uuid.UUID) (*magazine.Photograph, error) {
	Photo, err := u.repo.GetPhotoFromID(id)
	if err != nil {
		return nil, err
	}

	return Photo, nil
}

// Lists Photos by profile id
func (u PhotoService) ListPhotosByProfileId(id uuid.UUID) ([]*magazine.Photograph, error) {
	Photo, err := u.repo.GetPhotoFromCreatorId(id)
	if err != nil {
		return nil, err
	}

	return Photo, nil
}

// Lists Photos by type
func (u PhotoService) ListPhotosByType(story_type string) ([]*magazine.Photograph, error) {
	Photo, err := u.repo.GetPhotoFromType(story_type)
	if err != nil {
		return nil, err
	}

	return Photo, nil
}

// Update Photo by id in our database
func (u PhotoService) UpdatePhoto(id uuid.UUID, patch *map[string]interface{}) error {
	err := u.repo.PatchPhoto(id, patch)
	if err != nil {
		return err
	}

	return nil
}

// Delete Photo by in our database
func (u PhotoService) DeletePhoto(id uuid.UUID) error {
	err := u.repo.DeletePhoto(id)
	if err != nil {
		return err
	}

	return nil
}

// Permanent Delete Photo by in our database permanently
func (u PhotoService) PermanentDeletePhoto(id uuid.UUID) error {
	err := u.repo.PermanentDeletePhoto(id)
	if err != nil {
		return err
	}

	return nil
}

func (u PhotoService) BeforeCreate(Photo *magazine.Photograph) *magazine.Photograph {
	Photo.ID = uuid.New()
	create := time.Now()
	Photo.CreatedOn = &create
	Photo.UpdatedOn = &create

	return Photo
}
