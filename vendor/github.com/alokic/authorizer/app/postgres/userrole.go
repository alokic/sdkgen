package postgres

import (
	"context"
	"time"

	"github.com/alokic/authorizer/app"
	sql "github.com/alokic/gopkg/sql/v2"
	"github.com/alokic/gopkg/typeutils"
)

var (
	// UserRoleDeleteQuery deletes all roles for an existing user
	UserRoleDeleteQuery = "DELETE FROM user_roles WHERE user_id = $1"
)

// UserRole is postgres repo for users.
type UserRole struct {
	db        DB
	logger    Logger
	fieldInfo *sql.FieldInfo
}

type userRoleAttrib struct {
	UserID    uint64    `db:"user_id"`
	RoleID    uint64    `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// NewUserRole is is userrole constructor.
func NewUserRole(d DB, logger Logger) *UserRole {
	return &UserRole{db: d, logger: logger, fieldInfo: sql.GenFieldInfo(DriverName, userRoleAttrib{})}
}

// BatchUpdate will overwrite roles for a given user
func (repo *UserRole) BatchUpdate(ctx context.Context, urs []*app.UserRole) ([]*app.UserRole, error) {
	var a []*userRoleAttrib
	for _, ur := range urs {
		t := toUserRoleAttribs(ur)
		t.CreatedAt = time.Now().UTC()
		t.UpdatedAt = t.CreatedAt

		err := repo.validate(t)
		if err != nil {
			repo.logger.Log("method", "batch_update", "error", err)
			return nil, err
		}

		a = append(a, t)
	}

	repo.db.Transaction(func(tx *sql.Tx) (interface{}, error) {
		_, err := tx.Exec(UserRoleDeleteQuery, urs[0].UserId)
		if err != nil {
			return nil, err
		}

		var r []interface{}
		for _, ir := range a {
			r = append(r, ir)
		}

		stmt, params := sql.BatchInsertStatement("user_roles", r, repo.fieldInfo)
		return tx.Exec(stmt, params...)
	})

	var ura []*app.UserRole
	for _, ur := range a {
		ura = append(ura, ur.toUserRole())
	}
	return ura, nil
}

func toUserRoleAttribs(t *app.UserRole) *userRoleAttrib {
	e := &userRoleAttrib{}
	if t == nil {
		return e
	}
	e.UserID = t.UserId
	e.RoleID = t.RoleId
	return e
}

func (repo *UserRole) validate(p *userRoleAttrib, validateOnlyPresent ...bool) error {
	_, err := repo.fieldInfo.Validate(p, validateOnlyPresent...)

	if err == nil {
		return err
	}

	return nil
}

func (t *userRoleAttrib) toUserRole() *app.UserRole {
	e := &app.UserRole{}
	if t == nil {
		return e
	}

	e.UserId = t.UserID
	e.RoleId = t.RoleID
	e.CreatedAt = typeutils.ToUnixTime(t.CreatedAt)
	e.UpdatedAt = typeutils.ToUnixTime(t.UpdatedAt)

	return e
}
