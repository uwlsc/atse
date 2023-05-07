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

//Magazine management component structure
type IMagazineMgmtComp struct {
	infrastructure.Database
}

//New Magazine Management Component creates a new Magazine component
func NewMagazineComp(db infrastructure.Database, logger lib.Logger) IMagazineMgmtComp {
	return IMagazineMgmtComp{db}
}

//Creates Magazine in our database
func (a IMagazineMgmtComp) CreateMagazine(magazine magazine.Magazine) error {
	sql, args, err := sqrl.Insert("magazines").
		Columns("id", "creator_id", "magazine_code", "issue_code", "placement", "remarks").
		Values(magazine.ID, magazine.CreatedBy, magazine.MagazineCode, magazine.IssueCode, magazine.Placement, magazine.Remarks).
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

// Lists all the Magazines from our database
func (s IMagazineMgmtComp) ListMagazines() ([]*magazine.Magazine, error) {
	var stores []*magazine.Magazine
	sql, args, err := sqrl.Select("*").From("magazines").Where(sqrl.Eq{"deleted_on": nil}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), s, &stores, sql, args[:]...); err != nil {
		return nil, err
	}

	return stores, nil
}

//Get One Magazine from our database based on id
func (a IMagazineMgmtComp) GetMagazineFromID(id uuid.UUID) (*magazine.Magazine, error) {
	var Magazine magazine.Magazine

	sql, args, err := sqrl.Select("*").From("magazines").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), a, &Magazine, sql, args[:]...); err != nil {
		return nil, err
	}

	return &Magazine, nil
}

// Get Magazines from our database based on profile id
func (a IMagazineMgmtComp) GetMagazineFromCreatorId(id uuid.UUID) ([]*magazine.Magazine, error) {
	var magazine []*magazine.Magazine
	sql, args, err := sqrl.Select("*").From("magazines").Where(sqrl.Eq{"creator_id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &magazine, sql, args[:]...); err != nil {
		return nil, err
	}

	return magazine, nil
}

// Get Magazines from our database based on type
func (a IMagazineMgmtComp) GetMagazineFromType(magazine_code string) ([]*magazine.Magazine, error) {
	var magazine []*magazine.Magazine
	sql, args, err := sqrl.Select("*").From("magazines").Where(sqrl.Eq{"magazine_code": magazine_code}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &magazine, sql, args[:]...); err != nil {
		return nil, err
	}

	return magazine, nil
}

// Delete Magazine soft deletes the Magazine in our database
func (a IMagazineMgmtComp) DeleteMagazine(id uuid.UUID) error {
	sql, arg, err := sqrl.Update("magazines").SetMap(gin.H{"deleted_on": time.Now()}).
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

//Permanent Delete Magazine permanently deletes the Magazine in our database
func (a IMagazineMgmtComp) PermanentDeleteMagazine(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("magazines").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
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

//PatchMagazine updates the Magazine in our database
func (a IMagazineMgmtComp) PatchMagazine(id uuid.UUID, patch *map[string]interface{}) error {
	sql, args, err := sqrl.Update("magazines").SetMap(*patch).Where(sqrl.Eq{"id": id}).
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
