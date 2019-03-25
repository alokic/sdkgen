package postgres

import (
	"context"
	"time"

	"fmt"

	"github.com/alokic/authorizer/app"
	sql "github.com/alokic/gopkg/sql/v2"
	"github.com/alokic/gopkg/typeutils"
)

var (
	// PermissionUpdateQuery is query template for updating permissions.
	PermissionUpdateQuery = fmt.Sprintf("UPDATE permissions SET %s WHERE %s", "%s", "%s")
	// PermissionGetQuery is lightest query template to get a Permission.
	PermissionGetQuery = "SELECT * FROM permissions WHERE id = $1"
	// PermissionListAllQuery is query template to list all permissions.
	PermissionListAllQuery = "SELECT * FROM permissions"
	// PermissionListInIntervalQuery is query template to list all permissions.
	PermissionListInIntervalQuery = "SELECT * FROM permissions where updated_at >= $1"
	// PermissionStateQuery is query template to fetch state of permission along with service.
	PermissionStateQuery = "SELECT permissions.state as state, services.id as sid, services.state as sstate from permissions inner join services on permissions.service_id = services.id where permissions.id = $1"
	// PermissionListByAccessPolicyQuery is query template to list permissions by access query
	PermissionListByAccessPolicyQuery = `
		SELECT permissions.* from permissions INNER JOIN access_policies_permissions ON permissions.id = access_policies_permissions.permission_id WHERE access_policies_permissions.access_policy_id = $1;
	`
)

// Permission is postgres repo for permissions.
type Permission struct {
	db        DB
	logger    Logger
	fieldInfo *sql.FieldInfo
}

type permissionAttrib struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	ServiceID   uint64    `db:"service_id"`
	Path        string    `db:"path"`
	Action      string    `db:"action"`
	Description string    `db:"description"`
	State       string    `db:"state"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type permissionPermissionState struct {
	SID    uint64 `db:"sid"`
	State  string `db:"state"`
	SState string `db:"sstate"`
}
type permissionAttribs []*permissionAttrib

// NewPermission is the permission constructor.
func NewPermission(d DB, logger Logger) *Permission {
	return &Permission{db: d, logger: logger, fieldInfo: sql.GenFieldInfo(DriverName, permissionAttrib{})}
}

// Create a permission in the db.
func (repo *Permission) Create(ctx context.Context, e *app.Permission) (*app.Permission, error) {
	var p *permissionAttrib
	{
		p = toPermissionAttribs(e)

		p.CreatedAt = time.Now().UTC()
		p.UpdatedAt = p.CreatedAt

		err := repo.validate(p)
		if err != nil {
			return nil, err
		}
	}

	err := BatchInsert(repo.db, "permissions", []interface{}{p}, repo.fieldInfo)
	if err != nil {
		return nil, err
	}

	return p.toPermission(), nil
}

// Get a Permission.
func (repo *Permission) Get(ctx context.Context, id uint64) (*app.Permission, error) {
	var t []permissionAttrib

	if err := repo.db.Select(&t, PermissionGetQuery, id); err != nil {
		repo.logger.Log("method", "get", "error", err)
		return nil, err
	}

	if len(t) == 0 {
		return nil, nil
	}

	return t[0].toPermission(), nil
}

// Update a permission in the db
func (repo *Permission) Update(ctx context.Context, e *app.Permission) (*app.Permission, error) {
	p := toPermissionAttribs(e)

	if err := repo.update(ctx, p); err != nil {
		return nil, err
	}

	return p.toPermission(), nil
}

// GetState of a permission along with in the db
func (repo *Permission) GetState(ctx context.Context, id uint64) (*app.Permission, *app.Svc, error) {
	var t []permissionPermissionState

	if err := repo.db.Select(&t, PermissionStateQuery, id); err != nil {
		repo.logger.Log("method", "get_state", "error", err)
		return nil, nil, err
	}

	if len(t) == 0 {
		return nil, nil, nil
	}

	return &app.Permission{Id: id, State: t[0].State}, &app.Svc{Id: t[0].SID, State: t[0].SState}, nil
}

// Block a permission from the db.
func (repo *Permission) Block(ctx context.Context, id uint64) error {
	_, err := repo.Update(ctx, &app.Permission{Id: id, State: app.PermissionStateBlocked})
	if err != nil {
		repo.logger.Log("method", "block", "error", err)
	}
	return err
}

// List lists permissions from the db based on the input.
func (repo *Permission) List(ctx context.Context, apID uint64) ([]*app.Permission, error) {
	var pa []*app.Permission
	var err error

	{
		if apID == 0 {
			pa, err = repo.ListAll(ctx)
		} else {
			pa, err = repo.ListByAccessPolicy(ctx, apID)
		}
	}

	if err != nil {
		return nil, err
	}

	return pa, nil
}

// ListAll lists all permissions from the db.
func (repo *Permission) ListAll(ctx context.Context) ([]*app.Permission, error) {
	var pa []*permissionAttrib
	{
		if err := repo.db.Select(&pa, PermissionListAllQuery); err != nil {
			return nil, err
		}
	}
	var ps []*app.Permission
	{
		for _, a := range pa {
			ps = append(ps, a.toPermission())
		}
	}

	return ps, nil
}

// ListInInterval lists all permissions from the db in defined interval in seconds.
func (repo *Permission) ListInInterval(ctx context.Context, seconds int64) ([]*app.Permission, error) {
	var pa []*permissionAttrib
	{
		ts := time.Now().Add(-time.Duration(seconds) * time.Second).UTC()
		if err := repo.db.Select(&pa, PermissionListInIntervalQuery, ts); err != nil {
			return nil, err
		}
	}
	ps := permissionAttribs(pa).toPermissions()
	return ps, nil
}

// ListByAccessPolicy lists permissions by access policy from the db.
func (repo *Permission) ListByAccessPolicy(ctx context.Context, apID uint64) ([]*app.Permission, error) {
	var pa []*permissionAttrib
	{
		if err := repo.db.Select(&pa, PermissionListByAccessPolicyQuery, apID); err != nil {
			repo.logger.Log("method", "listByAccessPolicy", "error", err)
			return nil, err
		}
	}
	ps := permissionAttribs(pa).toPermissions()
	return ps, nil
}

func (pp permissionAttribs) toPermissions() []*app.Permission {
	var ps []*app.Permission
	{
		for _, a := range pp {
			ps = append(ps, a.toPermission())
		}
	}
	return ps
}

func (repo *Permission) update(ctx context.Context, p *permissionAttrib, args ...string) error {
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
		s, err := sql.PartialUpdateStmt(p, "permissions", condition, repo.fieldInfo)
		if err != nil {
			return err
		}
		setStatement = s
	}

	_, err := repo.db.Exec(setStatement)
	return err
}

func toPermissionAttribs(p *app.Permission) *permissionAttrib {
	pa := &permissionAttrib{}

	if p == nil {
		return pa
	}

	pa.ServiceID = p.ServiceId
	pa.Name = p.Name
	pa.Description = p.Description
	pa.Action = p.Action
	pa.State = p.State
	pa.ID = p.Id
	pa.Path = p.Path

	return pa
}

func (p *permissionAttrib) toPermission() *app.Permission {
	ap := &app.Permission{}
	if p == nil {
		return ap
	}

	ap.Id = p.ID
	ap.Name = p.Name
	ap.Action = p.Action
	ap.Description = p.Description
	ap.State = p.State
	ap.ServiceId = p.ServiceID
	ap.Path = p.Path
	ap.CreatedAt = typeutils.ToUnixTime(p.CreatedAt)
	ap.UpdatedAt = typeutils.ToUnixTime(p.UpdatedAt)

	return ap
}

func (repo *Permission) validate(p *permissionAttrib, validateOnlyPresent ...bool) error {
	_, err := repo.fieldInfo.Validate(p, validateOnlyPresent...)

	if err == nil {
		return err
	}

	return nil
}
