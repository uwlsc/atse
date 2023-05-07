package services

import (
	"magazine_api/component"
	"magazine_api/lib"
	"magazine_api/models/magazine"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AdvertService service layer
type AdvertService struct {
	logger lib.Logger
	repo   component.IAdMgmtComp
}

// NewAdvertService creates new instance of AdvertService
func NewAdvertService(logger lib.Logger, repo component.IAdMgmtComp) AdvertService {
	return AdvertService{logger: logger, repo: repo}
}

// Creates the Advert in database
func (u AdvertService) CreateAdvert(ad *magazine.Advert) (*magazine.Advert, error) {
	ad = u.BeforeCreate(ad)

	err := u.repo.CreateAd(*ad)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

// Lists stories form database
func (p AdvertService) ListAds(c *gin.Context) ([]*magazine.Advert, error) {
	stories, err := p.repo.ListAds()
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// Gets Advert by id from database
func (u AdvertService) GetAdvertById(id uuid.UUID) (*magazine.Advert, error) {
	Advert, err := u.repo.GetAdFromID(id)
	if err != nil {
		return nil, err
	}

	return Advert, nil
}

// Lists Adverts by profile id
func (u AdvertService) ListAdsByProfileId(id uuid.UUID) ([]*magazine.Advert, error) {
	Advert, err := u.repo.GetAdFromCreatorId(id)
	if err != nil {
		return nil, err
	}

	return Advert, nil
}

// Update Advert by id in our database
func (u AdvertService) UpdateAdvert(id uuid.UUID, patch *map[string]interface{}) error {
	err := u.repo.PatchAd(id, patch)
	if err != nil {
		return err
	}

	return nil
}

// Delete Advert by in our database
func (u AdvertService) DeleteAdvert(id uuid.UUID) error {
	err := u.repo.DeleteAd(id)
	if err != nil {
		return err
	}

	return nil
}

// Permanent Delete Advert by in our database permanently
func (u AdvertService) PermanentDeleteAdvert(id uuid.UUID) error {
	err := u.repo.PermanentDeleteAd(id)
	if err != nil {
		return err
	}

	return nil
}

func (u AdvertService) BeforeCreate(Advert *magazine.Advert) *magazine.Advert {
	Advert.ID = uuid.New()
	create := time.Now()
	Advert.CreatedOn = &create
	Advert.UpdatedOn = &create

	return Advert
}
