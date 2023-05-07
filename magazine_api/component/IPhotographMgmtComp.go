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

type IPhotographMgmtComp struct {
	infrastructure.Database
}

//New Photo Management Component creates a new Photo component
func NewPhotographComp(db infrastructure.Database, logger lib.Logger) IPhotographMgmtComp {
	return IPhotographMgmtComp{db}
}

//Creates Photo in our database
func (i IPhotographMgmtComp) CreatePhoto(photo magazine.Photograph) error {
	sql, args, err := sqrl.Insert("photographs").
		Columns("id", "photo_id", "photo_code", "creator_id", "photo_title", "photo_type", "url", "remarks").
		Values(photo.ID, photo.PhotographID, photo.PhotographCode, photo.CreatedBy, photo.PhotographTitle, photo.PhotographType, photo.DocumentURL, photo.Remarks).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}

	exec, err := i.Exec(context.Background(), sql, args[:]...)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return errors.New("not inserted")
	}

	return nil
}

// Lists all the Photographs from our database
func (s IPhotographMgmtComp) ListPhotos() ([]*magazine.Photograph, error) {
	var stores []*magazine.Photograph
	sql, args, err := sqrl.Select("*").From("photographs").Where(sqrl.Eq{"deleted_on": nil}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), s, &stores, sql, args[:]...); err != nil {
		return nil, err
	}

	return stores, nil
}

//Get One Photo from our database based on id
func (a IPhotographMgmtComp) GetPhotoFromID(id uuid.UUID) (*magazine.Photograph, error) {
	var Photo magazine.Photograph

	sql, args, err := sqrl.Select("*").From("photographs").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), a, &Photo, sql, args[:]...); err != nil {
		return nil, err
	}

	return &Photo, nil
}

// Get Photos from our database based on profile id
func (a IPhotographMgmtComp) GetPhotoFromCreatorId(id uuid.UUID) ([]*magazine.Photograph, error) {
	var Photo []*magazine.Photograph
	sql, args, err := sqrl.Select("*").From("photographs").Where(sqrl.Eq{"creator_id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &Photo, sql, args[:]...); err != nil {
		return nil, err
	}

	return Photo, nil
}

// Get Photos from our database based on type
func (a IPhotographMgmtComp) GetPhotoFromType(photo_type string) ([]*magazine.Photograph, error) {
	var Photo []*magazine.Photograph
	sql, args, err := sqrl.Select("*").From("photographs").Where(sqrl.Eq{"photo_type": photo_type}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &Photo, sql, args[:]...); err != nil {
		return nil, err
	}

	return Photo, nil
}

// Delete Photo soft deletes the Photo in our database
func (a IPhotographMgmtComp) DeletePhoto(id uuid.UUID) error {
	sql, arg, err := sqrl.Update("photographs").SetMap(gin.H{"deleted_on": time.Now()}).
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

//Permanent Delete Photo permanently deletes the Photo in our database
func (a IPhotographMgmtComp) PermanentDeletePhoto(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("photographs").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
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

//Patch Photo updates the Photo in our database
func (a IPhotographMgmtComp) PatchPhoto(id uuid.UUID, patch *map[string]interface{}) error {
	sql, args, err := sqrl.Update("photographs").SetMap(*patch).Where(sqrl.Eq{"id": id}).
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
