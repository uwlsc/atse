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

//Story management component structure
type IStoryMgmtComp struct {
	infrastructure.Database
}

//New Story Management Component creates a new Story component
func NewStoryComp(db infrastructure.Database, logger lib.Logger) IStoryMgmtComp {
	return IStoryMgmtComp{db}
}

//Creates Story in our database
func (a IStoryMgmtComp) CreateStory(story magazine.Story) error {
	sql, args, err := sqrl.Insert("stories").
		Columns("id", "creator_id", "story_title", "story_type", "story_content", "remarks").
		Values(story.ID, story.CreatorId, story.StoryTitle, story.StoryType, story.StoryContent, story.Remarks).
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

// Lists all the Stories from our database
func (s IStoryMgmtComp) ListStories() ([]*magazine.Story, error) {
	var stores []*magazine.Story
	sql, args, err := sqrl.Select("*").From("stories").Where(sqrl.Eq{"deleted_on": nil}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), s, &stores, sql, args[:]...); err != nil {
		return nil, err
	}

	return stores, nil
}

//Get One Story from our database based on id
func (a IStoryMgmtComp) GetStoryFromID(id uuid.UUID) (*magazine.Story, error) {
	var Story magazine.Story

	sql, args, err := sqrl.Select("*").From("stories").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), a, &Story, sql, args[:]...); err != nil {
		return nil, err
	}

	return &Story, nil
}

// Get Stories from our database based on profile id
func (a IStoryMgmtComp) GetStoryFromCreatorId(id uuid.UUID) ([]*magazine.Story, error) {
	var story []*magazine.Story
	sql, args, err := sqrl.Select("*").From("stories").Where(sqrl.Eq{"creator_id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &story, sql, args[:]...); err != nil {
		return nil, err
	}

	return story, nil
}

// Get Stories from our database based on type
func (a IStoryMgmtComp) GetStoryFromType(story_type string) ([]*magazine.Story, error) {
	var story []*magazine.Story
	sql, args, err := sqrl.Select("*").From("stories").Where(sqrl.Eq{"story_type": story_type}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &story, sql, args[:]...); err != nil {
		return nil, err
	}

	return story, nil
}

// Delete Story soft deletes the Story in our database
func (a IStoryMgmtComp) DeleteStory(id uuid.UUID) error {
	sql, arg, err := sqrl.Update("stories").SetMap(gin.H{"deleted_on": time.Now()}).
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

//Permanent Delete Story permanently deletes the Story in our database
func (a IStoryMgmtComp) PermanentDeleteStory(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("stories").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
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

//PatchStory updates the Story in our database
func (a IStoryMgmtComp) PatchStory(id uuid.UUID, patch *map[string]interface{}) error {
	sql, args, err := sqrl.Update("stories").SetMap(*patch).Where(sqrl.Eq{"id": id}).
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
