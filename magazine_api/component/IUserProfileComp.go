package component

import (
	"context"
	"errors"
	"magazine_api/infrastructure"
	"magazine_api/lib"
	"magazine_api/models"
	"time"

	"github.com/elgris/sqrl"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileComponent struct {
	infrastructure.Database
}

func NewUserProfileComponent(db infrastructure.Database, logger lib.Logger) UserProfileComponent {
	return UserProfileComponent{db}
}

func (u UserProfileComponent) CreateUserProfile(profile models.UserProfile) error {
	sql, args, err := sqrl.Insert("user_profiles").
		Columns("id", "user_id", "name", "email", "contact_number", "picture", "created_on").
		Values(profile.ID, profile.UserId, profile.Name, profile.Email, profile.ContactNumber, profile.Picture, profile.CreatedOn).
		PlaceholderFormat(sqrl.Dollar).ToSql()

	if err != nil {
		return err
	}

	exec, err := u.Exec(context.Background(), sql, args[:]...)

	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return errors.New("not inserted")
	}

	return nil
}

// Get profile assign from id
func (u UserProfileComponent) GetUserProfileFromID(
	id uuid.UUID,
) (*models.UserProfile, error) {
	var profile models.UserProfile

	sql, args, err := sqrl.Select("*").From("profileassigns").Where(sqrl.Eq{"id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), u, &profile, sql, args[:]...); err != nil {
		return nil, err
	}

	return &profile, nil
}

// Patch CuttingAssignment in our database
func (u UserProfileComponent) PatchUserProfile(
	id uuid.UUID,
	patch *map[string]interface{},
) error {
	sql, args, err := sqrl.Update("user_profiles").SetMap(*patch).Where(sqrl.Eq{"id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil
	}

	exec, err := u.Exec(context.Background(), sql, args[:]...)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return errors.New("not updated")
	}

	return nil
}

// Delete profile order soft deletes the user in our database
func (u UserProfileComponent) DeleteUserProfile(id uuid.UUID) error {
	sql, arg, err := sqrl.Update("user_profiles").SetMap(gin.H{"deleted_on": time.Now()}).
		Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return err
	}

	exec, err := u.Exec(context.Background(), sql, arg[:]...)
	if err != nil {
		return err
	}

	count := exec.RowsAffected()
	if count != 1 {
		return err
	}

	return nil
}

// Permanent Delete orderprofile permanently delets the user in our database
func (u UserProfileComponent) PermanentDeleteUserProfile(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("user_profiles").
		Where(sqrl.Eq{"id": id}).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	exec, err := u.Exec(context.Background(), sql, arg[:]...)
	if err != nil {
		return err
	}

	count := exec.RowsAffected()
	if count != 1 {
		return err
	}

	return nil
}
