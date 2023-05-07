package services

import (
	"magazine_api/component"
	"magazine_api/lib"
	"magazine_api/models/magazine"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MagazineService service layer
type MagazineService struct {
	logger lib.Logger
	repo   component.IMagazineMgmtComp
}

// NewMagazineService creates new instance of MagazineService
func NewMagazineService(logger lib.Logger, repo component.IMagazineMgmtComp) MagazineService {
	return MagazineService{logger: logger, repo: repo}
}

// Creates the Magazine in database
func (u MagazineService) CreateMagazine(story *magazine.Magazine) (*magazine.Magazine, error) {
	story = u.BeforeCreate(story)

	err := u.repo.CreateMagazine(*story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// Lists stories form database
func (p MagazineService) ListMagazines(c *gin.Context) ([]*magazine.Magazine, error) {
	stories, err := p.repo.ListMagazines()
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// Gets Magazine by id from database
func (u MagazineService) GetMagazineById(id uuid.UUID) (*magazine.Magazine, error) {
	Magazine, err := u.repo.GetMagazineFromID(id)
	if err != nil {
		return nil, err
	}

	return Magazine, nil
}

// Lists Magazines by profile id
func (u MagazineService) ListMagazinesByProfileId(id uuid.UUID) ([]*magazine.Magazine, error) {
	Magazine, err := u.repo.GetMagazineFromCreatorId(id)
	if err != nil {
		return nil, err
	}

	return Magazine, nil
}

// Lists Magazines by type
func (u MagazineService) ListMagazinesByType(magazine_code string) ([]*magazine.Magazine, error) {
	Magazine, err := u.repo.GetMagazineFromType(magazine_code)
	if err != nil {
		return nil, err
	}

	return Magazine, nil
}

// Update Magazine by id in our database
func (u MagazineService) UpdateMagazine(id uuid.UUID, patch *map[string]interface{}) error {
	err := u.repo.PatchMagazine(id, patch)
	if err != nil {
		return err
	}

	return nil
}

// Delete Magazine by in our database
func (u MagazineService) DeleteMagazine(id uuid.UUID) error {
	err := u.repo.DeleteMagazine(id)
	if err != nil {
		return err
	}

	return nil
}

// Permanent Delete Magazine by in our database permanently
func (u MagazineService) PermanentDeleteMagazine(id uuid.UUID) error {
	err := u.repo.PermanentDeleteMagazine(id)
	if err != nil {
		return err
	}

	return nil
}

func (u MagazineService) BeforeCreate(Magazine *magazine.Magazine) *magazine.Magazine {
	Magazine.ID = uuid.New()
	create := time.Now()
	Magazine.CreatedOn = &create
	Magazine.UpdatedOn = &create

	return Magazine
}
