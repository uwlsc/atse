package component

import (
	"context"
	"errors"
	"magazine_api/api/serializers/responses"
	"magazine_api/infrastructure"
	"magazine_api/lib"
	"magazine_api/models"
	"time"

	"github.com/elgris/sqrl"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserComponent database structure
type UserComponent struct {
	infrastructure.Database
}

// NewUserComponent creates a new user component
func NewUserComponent(db infrastructure.Database, logger lib.Logger) UserComponent {
	return UserComponent{db}
}

// Creates user in our database
func (u UserComponent) CreateUser(user models.User) error {
	sql, args, err := sqrl.Insert("users").
		Columns("id", "role", "name", "email", "contact_number", "created_on").
		Values(user.ID, user.Role, user.Name, user.Email, user.ContactNumber, user.CreatedOn).
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

// Lists Deleted users from our database
func (u UserComponent) ListDeletedUsers() ([]*models.User, error) {
	var users []*models.User

	sql, args, err := sqrl.Select("*").From("users").Where(sqrl.NotEq{"deleted_on": nil}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), u, &users, sql, args[:]...); err != nil {
		return nil, err
	}

	return users, nil
}

// Lists users from our database
func (u UserComponent) ListEmployees(limit int64, offset int64) (map[string]interface{}, error) {
	var users []*responses.EmployeeSmall
	var count int64

	sql, args, err := sqrl.
		Select("employee_profile_id", "user_id", "name", "email", "role", "contact_number", "company_id").
		From("employee_view").
		Where(sqrl.Eq{"deleted_on": nil}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), u, &users, sql, args...); err != nil {
		return nil, err
	}

	sql, args, err = sqrl.
		Select("COUNT(*) AS total").
		From("employee_view").
		Where(sqrl.Eq{"deleted_on": nil}).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), u, &count, sql, args...); err != nil {
		return nil, err
	}

	return gin.H{"data": users, "count": count}, nil
}

// Lists users by type from our database
func (u UserComponent) ListEmployeesByType(limit int64, offset int64, userType string) ([]*responses.EmployeeSmall, error) {
	var users []*responses.EmployeeSmall
	sql, args, err := sqrl.
		Select("employee_profile_id", "user_id", "name", "email", "role", "contact_number", "company_id").
		From("employee_view").
		Where(sqrl.Eq{"deleted_on": nil}).
		Where(sqrl.Expr("? = ANY (role)", userType)).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(context.Background(), u, &users, sql, args...); err != nil {
		return nil, err
	}

	return users, nil
}

// Get One user from our database based on id
func (u UserComponent) GetProfileFromID(id uuid.UUID) (*responses.EmployeeAll, error) {
	var user responses.EmployeeAll

	sql, args, err := sqrl.
		Select("*").
		From("employee_view").
		Where(sqrl.Eq{"employee_profile_id": id}).
		PlaceholderFormat(sqrl.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), u, &user, sql, args[:]...); err != nil {
		return nil, err
	}

	return &user, nil
}

// Get One user from our database based on email
func (u UserComponent) GetUserFromEmail(email string) (*models.User, error) {
	var user models.User

	sql, args, err := sqrl.Select("*").From("users").Where(sqrl.Eq{"email": email}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), u, &user, sql, args[:]...); err != nil {
		return nil, err
	}

	return &user, nil
}

// Get One user from our database based on contact number
func (u UserComponent) GetUserFromContactNumber(contactNumber string) (*models.User, error) {
	var user models.User

	sql, args, err := sqrl.Select("*").From("users").Where(sqrl.Eq{"contact_number": contactNumber}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), u, &user, sql, args[:]...); err != nil {
		return nil, err
	}

	return &user, nil
}

// Get One user from our database based on id
func (u UserComponent) GetUserFromID(id uuid.UUID) (*models.User, error) {
	var user models.User

	sql, args, err := sqrl.Select("*").From("users").Where(sqrl.Eq{"id": id}).
		PlaceholderFormat(sqrl.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Get(context.Background(), u, &user, sql, args[:]...); err != nil {
		return nil, err
	}

	return &user, nil
}

// Patch Users updates the user in our database
func (u UserComponent) PatchUser(id uuid.UUID, patch *map[string]interface{}) error {
	sql, args, err := sqrl.Update("users").SetMap(*patch).Where(sqrl.Eq{"id": id}).
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

// Delete Users soft delets the user in our database
func (u UserComponent) DeleteUser(id uuid.UUID) error {
	sql, arg, err := sqrl.Update("users").SetMap(gin.H{"deleted_on": time.Now()}).
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

// Permanent Delete Users permanently delets the user in our database
func (u UserComponent) PermanentDeleteUser(id uuid.UUID) error {
	sql, arg, err := sqrl.Delete("users").Where(sqrl.Eq{"id": id}).PlaceholderFormat(sqrl.Dollar).ToSql()
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
