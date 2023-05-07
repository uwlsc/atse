package services

import (
	"magazine_api/component"
	"magazine_api/lib"
	"magazine_api/models/magazine"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MagazineIssueService service layer
type MagazineIssueService struct {
	logger lib.Logger
	comp   component.IIssueMgmtComp
}

// NewMagazineIssueService creates new instance of MagazineIssueService
func NewMagazineIssueService(logger lib.Logger, comp component.IIssueMgmtComp) MagazineIssueService {
	return MagazineIssueService{logger: logger, comp: comp}
}

// Creates the MagazineIssue in database
func (u MagazineIssueService) CreateMagazineIssue(issue *magazine.MagazineIssue) (*magazine.MagazineIssue, error) {
	issue = u.BeforeCreate(issue)

	err := u.comp.CreateIssue(*issue)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// Lists stories form database
func (p MagazineIssueService) ListIssues(c *gin.Context) ([]*magazine.MagazineIssue, error) {
	stories, err := p.comp.ListIssues()
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// Gets MagazineIssue by id from database
func (u MagazineIssueService) GetMagazineIssueById(id uuid.UUID) (*magazine.MagazineIssue, error) {
	MagazineIssue, err := u.comp.GetIssueFromID(id)
	if err != nil {
		return nil, err
	}

	return MagazineIssue, nil
}

// Lists MagazineIssues by profile id
func (u MagazineIssueService) ListIssuesByProfileId(id uuid.UUID) ([]*magazine.MagazineIssue, error) {
	MagazineIssue, err := u.comp.GetIssueFromCreatorId(id)
	if err != nil {
		return nil, err
	}

	return MagazineIssue, nil
}

// Lists Issues by type
func (u MagazineIssueService) ListIssuesByType(issue_type string) ([]*magazine.MagazineIssue, error) {
	MagazineIssue, err := u.comp.GetIssueFromType(issue_type)
	if err != nil {
		return nil, err
	}

	return MagazineIssue, nil
}

// Update MagazineIssue by id in our database
func (u MagazineIssueService) UpdateMagazineIssue(id uuid.UUID, patch *map[string]interface{}) error {
	err := u.comp.PatchIssue(id, patch)
	if err != nil {
		return err
	}

	return nil
}

// Delete MagazineIssue by in our database
func (u MagazineIssueService) DeleteMagazineIssue(id uuid.UUID) error {
	err := u.comp.DeleteIssue(id)
	if err != nil {
		return err
	}

	return nil
}

// Permanent Delete MagazineIssue by in our database permanently
func (u MagazineIssueService) PermanentDeleteMagazineIssue(id uuid.UUID) error {
	err := u.comp.PermanentDeleteIssue(id)
	if err != nil {
		return err
	}

	return nil
}

func (u MagazineIssueService) BeforeCreate(MagazineIssue *magazine.MagazineIssue) *magazine.MagazineIssue {
	MagazineIssue.ID = uuid.New()
	create := time.Now()
	MagazineIssue.CreatedOn = &create
	MagazineIssue.UpdatedOn = &create

	return MagazineIssue
}
