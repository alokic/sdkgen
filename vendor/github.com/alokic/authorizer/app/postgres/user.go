package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/alokic/authorizer/app"
	sql "github.com/alokic/gopkg/sql/v2"
	"github.com/alokic/gopkg/typeutils"
)

var (
	// UserListQuery is list query template for listing users.
	UserListQuery = "SELECT * FROM users where id IN (?)"
	// UserGetQuery is get query template to get an user.
	UserGetQuery = "SELECT * FROM users WHERE id = $1"
	// UserGetExternalIDQuery is get query template to get an user.
	UserGetExternalIDQuery = "SELECT * FROM users WHERE external_id = $1"
	// UserGetByEmailQuery is get by email query template to get an user.
	UserGetByEmailQuery = "SELECT * FROM users WHERE email = $1"
	// UserUpdateQuery is update query template to update an user.
	UserUpdateQuery = fmt.Sprintf("UPDATE users SET %s WHERE %s", "%s", "%s")
)

// User is postgres repo for users.
type User struct {
	db        DB
	logger    Logger
	fieldInfo *sql.FieldInfo
}

type userAttrib struct {
	ID         uint64    `db:"id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	ExternalID uint64    `db:"external_id"`
	State      string    `db:"state"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
type uas []userAttrib

// NewUser is is user constructor.
func NewUser(d DB, logger Logger) *User {
	return &User{db: d, logger: logger, fieldInfo: sql.GenFieldInfo(DriverName, userAttrib{})}
}

// Create a user in db.
func (repo *User) Create(ctx context.Context, e *app.User) (*app.User, error) {
	var t *userAttrib
	{
		t = toUserAtrribs(e)
		t.CreatedAt = time.Now().Round(time.Millisecond).UTC()
		t.UpdatedAt = t.CreatedAt

		err := repo.validate(t)
		if err != nil {
			repo.logger.Log("method", "create", "error", err)
			return nil, err
		}
	}

	err := BatchInsert(repo.db, "users", []interface{}{t}, repo.fieldInfo)
	if err != nil {
		repo.logger.Log("method", "create", "error", err)
		return nil, err
	}
	return t.toUser(), nil
}

// Get an user.
func (repo *User) Get(ctx context.Context, id uint64) (*app.User, error) {
	var t []userAttrib

	if err := repo.db.Select(&t, UserGetQuery, id); err != nil {
		repo.logger.Log("method", "get_by_email", "error", err)
		return nil, err
	}

	if len(t) == 0 {
		return nil, nil
	}
	return t[0].toUser(), nil
}

// GetByExternalID a user by external id.
func (repo *User) GetByExternalID(ctx context.Context, externalID uint64) (*app.User, error) {
	var t []userAttrib

	if err := repo.db.Select(&t, UserGetExternalIDQuery, externalID); err != nil {
		repo.logger.Log("method", "get_by_external_id", "error", err)
		return nil, err
	}

	if len(t) == 0 {
		return nil, nil
	}
	return t[0].toUser(), nil
}

// GetByEmail gets an user by email.
func (repo *User) GetByEmail(ctx context.Context, email string) (*app.User, error) {
	var t []userAttrib

	if err := repo.db.Select(&t, UserGetByEmailQuery, email); err != nil {
		repo.logger.Log("method", "get_by_email", "error", err)
		return nil, err
	}

	if len(t) == 0 {
		return nil, nil
	}
	return t[0].toUser(), nil
}

// List users.
func (repo *User) List(ctx context.Context, ids []uint64) ([]*app.User, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	query, args, err := sql.In(UserListQuery, ids)
	if err != nil {
		repo.logger.Log("method", "list", "error", err)
		return nil, err
	}

	query = sql.Rebind(DriverName, query)
	var ts []*userAttrib
	{
		if err := repo.db.Select(&ts, query, args...); err != nil {
			repo.logger.Log("method", "list", "error", err)
			return nil, err
		}
	}

	var e []*app.User
	{
		for _, v := range ts {
			e = append(e, v.toUser())
		}
	}
	return e, nil
}

// Update an user.
func (repo *User) Update(ctx context.Context, u *app.User) error {
	userAttrib := toUserAtrribs(u)

	if err := repo.update(ctx, userAttrib); err != nil {
		repo.logger.Log("method", "update", "error", err)
		return err
	}

	return nil
}

//
func (repo *User) update(ctx context.Context, p *userAttrib, args ...string) error {
	p.UpdatedAt = time.Now()

	var condition string
	{
		if len(args) == 0 {
			condition = fmt.Sprintf("id = %d", p.ID)
		} else {
			condition = args[0]
		}
	}

	var setStatement string
	{
		s, err := sql.PartialUpdateStmt(p, "users", condition, repo.fieldInfo)
		if err != nil {
			return err
		}
		setStatement = s
	}

	_, err := repo.db.Exec(setStatement)
	return err
}

func (t *userAttrib) toUser() *app.User {
	e := &app.User{}
	if t == nil {
		return e
	}

	e.Id = t.ID
	e.Name = t.Name
	e.Email = t.Email
	e.ExternalId = t.ExternalID
	e.State = t.State
	e.CreatedAt = typeutils.ToUnixTime(t.CreatedAt)
	e.UpdatedAt = typeutils.ToUnixTime(t.UpdatedAt)

	return e
}

// toUserAtrribs entity gateway
func toUserAtrribs(t *app.User) *userAttrib {
	e := &userAttrib{}
	if t == nil {
		return e
	}

	e.ID = t.Id
	e.Name = t.Name
	e.Email = t.Email
	e.ExternalID = t.ExternalId
	e.State = t.State

	return e
}

func (repo *User) validate(p *userAttrib, validateOnlyPresent ...bool) error {
	_, err := repo.fieldInfo.Validate(p, validateOnlyPresent...)

	if err == nil {
		return err
	}

	return nil
}
