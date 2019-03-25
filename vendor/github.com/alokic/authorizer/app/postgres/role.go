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
	// RoleListQuery is list query template for listing roles.
	RoleListQuery = "SELECT * FROM roles where id IN (?)"
	// RoleGetQuery is lightest query template to get a role.
	RoleGetQuery = "SELECT * FROM roles WHERE id = $1"
	// UserRoleListQuery gets roles for a given user
	UserRoleListQuery = "SELECT roles.* from roles INNER JOIN user_roles ON roles.id = user_roles.role_id WHERE user_roles.user_id = $1 and roles.state = 'active';"
)

// Role is postgres repo for roles.
type Role struct {
	db        DB
	logger    Logger
	fieldInfo *sql.FieldInfo
}

type roleAttrib struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	ServiceID   uint64    `db:"service_id"`
	Description string    `db:"description"`
	State       string    `db:"state"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// NewRole is role constructor.
func NewRole(d DB, logger Logger) *Role {
	return &Role{db: d, logger: logger, fieldInfo: sql.GenFieldInfo(DriverName, roleAttrib{})}
}

// Create a role in db.
func (repo *Role) Create(ctx context.Context, e *app.Role) (*app.Role, error) {
	var t *roleAttrib
	{
		t = toRoleAttribs(e)
		t.CreatedAt = time.Now().UTC()
		t.UpdatedAt = t.CreatedAt

		err := repo.validate(t)
		if err != nil {
			repo.logger.Log("method", "create", "error", err)
			return nil, err
		}
	}

	err := BatchInsert(repo.db, "roles", []interface{}{t}, repo.fieldInfo)
	if err != nil {
		repo.logger.Log("method", "create", "error", err)
		return nil, err
	}

	return t.toRole(), nil
}

// List roles.
func (repo *Role) List(ctx context.Context, ids []uint64, userid uint64) ([]*app.Role, error) {
	var err error
	var rs []*app.Role

	if len(ids) > 0 {
		rs, err = repo.ListRoles(ctx, ids)
	} else {
		rs, err = repo.ListUserRoles(ctx, userid)
	}

	return rs, err
}

// Get a role.
func (repo *Role) Get(ctx context.Context, id uint64) (*app.Role, error) {
	var t []roleAttrib

	if err := repo.db.Select(&t, RoleGetQuery, id); err != nil {
		repo.logger.Log("method", "get", "error", err)
		return nil, err
	}

	if len(t) == 0 {
		return nil, nil
	}

	return t[0].toRole(), nil
}

// ListRoles gets a list of multiple roles.
func (repo *Role) ListRoles(ctx context.Context, ids []uint64) ([]*app.Role, error) {
	query, args, err := sql.In(RoleListQuery, ids)
	if err != nil {
		repo.logger.Log("method", "list", "error", err)
		return nil, err
	}
	query = sql.Rebind(DriverName, query)
	var ts []*roleAttrib
	{
		if err := repo.db.Select(&ts, query, args...); err != nil {
			repo.logger.Log("method", "list", "error", err)
			return nil, err
		}
	}
	var rs []*app.Role
	{
		for _, v := range ts {
			rs = append(rs, v.toRole())
		}
	}

	return rs, nil
}

// ListUserRoles gets a list of roles for a user.
func (repo *Role) ListUserRoles(ctx context.Context, id uint64) ([]*app.Role, error) {
	query, args, err := sql.In(UserRoleListQuery, id)
	if err != nil {
		repo.logger.Log("method", "list", "error", err)
		return nil, err
	}
	var ura []*roleAttrib
	{
		if err := repo.db.Select(&ura, query, args...); err != nil {
			repo.logger.Log("method", "list", "error", err)
			return nil, err
		}
	}
	var urs []*app.Role
	{
		for _, v := range ura {
			urs = append(urs, v.toRole())
		}
	}
	return urs, nil
}

// Update is to update a given access policy.
func (repo *Role) Update(ctx context.Context, ap *app.Role) (*app.Role, error) {
	a := toRoleAttribs(ap)
	if err := repo.update(ctx, a); err != nil {
		repo.logger.Log("method", "update", "error", err)
		return nil, err
	}

	return a.toRole(), nil
}

func (repo *Role) update(ctx context.Context, a *roleAttrib, args ...string) error {
	a.UpdatedAt = time.Now()
	var condition string
	{
		if len(args) == 0 {
			condition = fmt.Sprintf("id = %d", a.ID)
		} else {
			condition = args[0]
		}
	}

	var setStatement string
	{
		s, err := sql.PartialUpdateStmt(a, "roles", condition, repo.fieldInfo)

		if err != nil {
			return err
		}

		setStatement = s
	}

	_, err := repo.db.Exec(setStatement)
	return err
}

func toRoleAttribs(t *app.Role) *roleAttrib {
	e := &roleAttrib{}
	if t == nil {
		return e
	}
	e.ID = t.Id
	e.Name = t.Name
	e.ServiceID = t.ServiceId
	e.Description = t.Description
	e.State = t.State

	return e
}

func (repo *Role) validate(p *roleAttrib, validateOnlyPresent ...bool) error {
	_, err := repo.fieldInfo.Validate(p, validateOnlyPresent...)

	if err == nil {
		return err
	}

	return nil
}

func (t *roleAttrib) toRole() *app.Role {
	e := &app.Role{}
	if t == nil {
		return e
	}

	e.Id = t.ID
	e.Name = t.Name
	e.ServiceId = t.ServiceID
	e.Description = t.Description
	e.State = t.State
	e.CreatedAt = typeutils.ToUnixTime(t.CreatedAt)
	e.UpdatedAt = typeutils.ToUnixTime(t.UpdatedAt)

	return e
}
