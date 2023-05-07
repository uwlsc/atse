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

//Issue management component structure
type IIssueMgmtComp struct {
	infrastructure.Database
}

//New Issue Management Component creates a new Issue component
func NewIssueComp(db infrastructure.Database, logger lib.Logger) IIssueMgmtComp {
	return IIssueMgmtComp{db}
}

//Creates Issue in our database
func (a IIssueMgmtComp) CreateIssue(issue magazine.MagazineIssue) error {
	sql, args, err := sqrl.Insert("magazine_issues").
		Columns("id", "issue_code", "content_code", "advert_code", "remarks").
		Values(issue.ID, issue.IssueCode, issue.ContentCode, issue.AdvertCode, issue.Remarks).
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

// Lists all the Issues from our database
func (s IIssueMgmtComp) ListIssues() ([]*magazine.MagazineIssue, error) {
	var stores []*magazine.MagazineIssue
	sql, args, err := sqrl.Select("*").From("magazine_issues").Where(sqrl.Eq{"deleted_on": nil}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), s, &stores, sql, args[:]...); err != nil {
		return nil, err
	}

	return stores, nil
}

//Get One Issue from our database based on id
func (a IIssueMgmtComp) GetIssueFromID(id uuid.UUID) (*magazine.MagazineIssue, error) {
	var Issue magazine.MagazineIssue

	sql, args, err := sqrl.Select("*").From("magazine_issues").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), a, &Issue, sql, args[:]...); err != nil {
		return nil, err
	}

	return &Issue, nil
}

// Get Issues from our database based on profile id
func (a IIssueMgmtComp) GetIssueFromCreatorId(id uuid.UUID) ([]*magazine.MagazineIssue, error) {
	var issue []*magazine.MagazineIssue
	sql, args, err := sqrl.Select("*").From("magazine_issues").Where(sqrl.Eq{"creator_id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &issue, sql, args[:]...); err != nil {
		return nil, err
	}

	return issue, nil
}

// Get Issues from our database based on type
func (a IIssueMgmtComp) GetIssueFromType(issue_type string) ([]*magazine.MagazineIssue, error) {
	var issue []*magazine.MagazineIssue
	sql, args, err := sqrl.Select("*").From("magazine_issues").Where(sqrl.Eq{"issue_type": issue_type}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), a, &issue, sql, args[:]...); err != nil {
		return nil, err
	}

	return issue, nil
}

// Delete Issue soft deletes the Issue in our database
func (a IIssueMgmtComp) DeleteIssue(id uuid.UUID) error {
	sql, arg, err := sqrl.Update("magazine_issues").SetMap(gin.H{"deleted_on": time.Now()}).
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

//Permanent Delete Issue permanently deletes the Issue in our database
func (a IIssueMgmtComp) PermanentDeleteIssue(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("magazine_issues").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
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

//PatchIssue updates the Issue in our database
func (a IIssueMgmtComp) PatchIssue(id uuid.UUID, patch *map[string]interface{}) error {
	sql, args, err := sqrl.Update("magazine_issues").SetMap(*patch).Where(sqrl.Eq{"id": id}).
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
