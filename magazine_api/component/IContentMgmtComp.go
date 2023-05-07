package component

import (
	"context"
	"errors"
	"magazine_api/infrastructure"
	"magazine_api/lib"
	"magazine_api/models/magazine"
	"time"

	"github.com/elgris/sqrl"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Content management component structure
type IContentMgmtComp struct {
	infrastructure.Database
}

//New Content Management Component creates a new Content component
func NewContentComp(db infrastructure.Database, logger lib.Logger) IContentMgmtComp {
	return IContentMgmtComp{db}
}

//Creates Content in our database
func (a IContentMgmtComp) CreateContent(content magazine.Content) error {
	sql, args, err := sqrl.Insert("contents").
		Columns("id", "content_code", "story_code", "photo_code", "created_on", "created_by", "creator_name", "remarks").
		Values(content.ID, content.ContentCode, content.StoryCode, content.PhotographCode, content.CreatedBy, content.CreatorName, content.Remarks).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}

	exec, err := a.Exec(context.Background(), sql, args[:]...)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return errors.New("not inserted")
	}

	return nil
}

// Lists all the Contents from our database
func (s IContentMgmtComp) ListContents() ([]*magazine.Content, error) {
	var stores []*magazine.Content
	sql, args, err := sqrl.Select("*").From("contents").Where(sqrl.Eq{"deleted_on": nil}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), s, &stores, sql, args[:]...); err != nil {
		return nil, err
	}

	return stores, nil
}

//Get One Content from our database based on id
func (a IContentMgmtComp) GetContentFromID(id uuid.UUID) (*magazine.Content, error) {
	var Content magazine.Content

	sql, args, err := sqrl.Select("*").From("contents").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), a, &Content, sql, args[:]...); err != nil {
		return nil, err
	}

	return &Content, nil
}

// Get Contents from our database based on profile id
func (a IContentMgmtComp) GetContentFromCreatorId(id uuid.UUID) ([]*magazine.Content, error) {
	var Content []*magazine.Content
	sql, args, err := sqrl.Select("*").From("contents").Where(sqrl.Eq{"creator_id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &Content, sql, args[:]...); err != nil {
		return nil, err
	}

	return Content, nil
}

// Get Contents from our database based on type
func (a IContentMgmtComp) GetContentFromType(story_type string) ([]*magazine.Content, error) {
	var Content []*magazine.Content
	sql, args, err := sqrl.Select("*").From("contents").Where(sqrl.Eq{"story_type": story_type}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &Content, sql, args[:]...); err != nil {
		return nil, err
	}

	return Content, nil
}

// Delete Content soft deletes the Content in our database
func (a IContentMgmtComp) DeleteContent(id uuid.UUID) error {
	sql, arg, err := sqrl.Update("contents").SetMap(gin.H{"deleted_on": time.Now()}).
		Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}

	exec, err := a.Exec(context.Background(), sql, arg[:]...)
	if err != nil {
		return err
	}

	count := exec.RowsAffected()
	if count != 1 {
		return err
	}

	return nil
}

//Permanent Delete Content permanently deletes the Content in our database
func (a IContentMgmtComp) PermanentDeleteContent(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("contents").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}

	exec, err := a.Exec(context.Background(), sql, arg[:]...)
	if err != nil {
		return err
	}

	count := exec.RowsAffected()
	if count != 1 {
		return err
	}

	return nil
}

//PatchContent updates the Content in our database
func (a IContentMgmtComp) PatchContent(id uuid.UUID, patch *map[string]interface{}) error {
	sql, args, err := sqrl.Update("contents").SetMap(*patch).Where(sqrl.Eq{"id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()

	if err != nil {
		return nil
	}

	exec, err := a.Exec(context.Background(), sql, args[:]...)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return errors.New("not updated")
	}

	return nil
}
