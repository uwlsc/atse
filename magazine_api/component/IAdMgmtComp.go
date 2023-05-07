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

//Advertisement Management Component structure
type IAdMgmtComp struct {
	infrastructure.Database
}

//NewIAdMgmtComp create new Advertisement Management component
func NewAdMgmtComp(db infrastructure.Database, logger lib.Logger) IAdMgmtComp {
	return IAdMgmtComp{db}
}

//Create Advert in our Database
func (i IAdMgmtComp) CreateAd(ad magazine.Advert) error {
	sql, args, err := sqrl.Insert("adverts").
		Columns("id", "advert_code", "advert_title", "advert_content", "advert_type", "url", "created_on", "created_by", "creator_name").
		Values(ad.ID, ad.AdvertCode, ad.AdvertTitle, ad.AdvertContent, ad.AdvertType, ad.AdvertURL, ad.CreatedOn, ad.CreatedBy, ad.CreatorName).
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

//Lists deleted adverts from database
func (i IAdMgmtComp) ListDeletedAd() ([]*magazine.Advert, error) {
	var ads []*magazine.Advert
	sql, args, err := sqrl.Select("*").From("adverts").Where(sqrl.NotEq{"deleted_on": nil}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), i, &ads, sql, args[:]...); err != nil {
		return nil, err
	}

	return ads, nil

}

//Lists adverts from database
func (i IAdMgmtComp) ListAds() ([]*magazine.Advert, error) {
	var ads []*magazine.Advert
	sql, args, err := sqrl.Select("*").From("adverts").Where(sqrl.Eq{"deleted_on": nil}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), i, &ads, sql, args[:]...); err != nil {
		return nil, err
	}

	return ads, nil

}

// Get Adverts from our database based on profile id
func (a IAdMgmtComp) GetAdFromCreatorId(id uuid.UUID) ([]*magazine.Advert, error) {
	var ad []*magazine.Advert
	sql, args, err := sqrl.Select("*").From("adverts").Where(sqrl.Eq{"creator_id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &ad, sql, args[:]...); err != nil {
		return nil, err
	}

	return ad, nil
}

//Gets single ad from datbase using ID
func (i IAdMgmtComp) GetAdFromID(id uuid.UUID) (*magazine.Advert, error) {
	var ad magazine.Advert

	sql, args, err := sqrl.Select("*").From("adverts").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), i, &ad, sql, args[:]...); err != nil {
		return nil, err
	}

	return &ad, nil
}

//Update advert in  our database
func (i IAdMgmtComp) PatchAd(id uuid.UUID, patch *map[string]interface{}) error {

	sql, args, err := sqrl.Update("adverts").SetMap(*patch).Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil
	}

	exec, err := i.Exec(context.Background(), sql, args[:]...)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return errors.New("not updated")
	}

	return nil
}

// Delete Advert in our database
func (i IAdMgmtComp) DeleteAd(id uuid.UUID) error {

	sql, arg, err := sqrl.Update("adverts").SetMap(gin.H{"deleted_on": time.Now()}).
		Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}

	exec, err := i.Exec(context.Background(), sql, arg[:]...)
	if err != nil {
		return err
	}

	count := exec.RowsAffected()
	if count != 1 {
		return err
	}

	return nil
}

func (i IAdMgmtComp) PermanentDeleteAd(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("adverts").
		Where(sqrl.Eq{"id": id}).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	exec, err := i.Exec(context.Background(), sql, arg[:]...)
	if err != nil {
		return err
	}

	count := exec.RowsAffected()
	if count != 1 {
		return err
	}

	return nil
}
