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
	// AccessPolicyUpdateQuery is query template for updating access policies.
	AccessPolicyUpdateQuery = fmt.Sprintf("UPDATE access_policies SET %s WHERE %s", "%s", "%s")
	// AccessPolicyGetQuery is lightest query template to get a AccessPolicy.
	AccessPolicyGetQuery = "SELECT * FROM access_policies WHERE id = $1"
	// AccessPolicyListQuery is query template for listing access policies.
	AccessPolicyListQuery = "SELECT * FROM access_policies where id IN (?)"
	// AccessPolicyDeleteQuery is query template to delete access policies
	AccessPolicyDeleteQuery = "DELETE FROM access_policies WHERE id = $1"
	// AccessPolicyListByRoleIDQuery is query template to get access policies by user id.
	AccessPolicyListByRoleIDQuery = "SELECT access_policies.* from access_policies INNER JOIN role_access_policies ON access_policies.id = role_access_policies.access_policy_id AND role_access_policies.role_id = $1"
	// AccessPolicyListByUserPermissionQuery is query template to get list of acccess_policy_id for a permission per user.
	AccessPolicyListByUserPermissionQuery = `
	select access_policies_permissions.access_policy_id from access_policies_permissions where permission_id = $1 and access_policy_id IN (
        select DISTINCT id as access_policy_id from access_policies
        INNER JOIN (
			select access_policy_id from role_access_policies where role_access_policies.role_id IN (
				select user_roles.role_id from user_roles where user_id = $2
			)
		) user_roles_access_polices 
		ON access_policies.id = user_roles_access_polices.access_policy_id 
		AND access_policies.state = 'active'
    );
	`
)

// AccessPolicy is postgres repo for access policies.
type AccessPolicy struct {
	db        DB
	logger    Logger
	fieldInfo *sql.FieldInfo
}

type accessPolicyAttrib struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	State     string    `db:"state"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// NewAccessPolicy is an access policy constructor.
func NewAccessPolicy(d DB, logger Logger) *AccessPolicy {
	return &AccessPolicy{db: d, logger: logger, fieldInfo: sql.GenFieldInfo(DriverName, accessPolicyAttrib{})}
}

// Create will create access policies.
func (repo *AccessPolicy) Create(ctx context.Context, ap *app.AccessPolicy) (*app.AccessPolicy, error) {
	var a *accessPolicyAttrib
	{
		a = toAccessPolicyAttribs(ap)
		a.CreatedAt = time.Now().UTC()
		a.UpdatedAt = a.CreatedAt

		err := repo.validate(a)

		if err != nil {
			repo.logger.Log("method", "create", "error", err)
			return nil, err
		}
	}

	err := BatchInsert(repo.db, "access_policies", []interface{}{a}, repo.fieldInfo)
	if err != nil {
		return nil, err
	}

	return a.toAccessPolicy(), nil
}

// List is to list access policies.
func (repo *AccessPolicy) List(ctx context.Context, ids []uint64, rID uint64) ([]*app.AccessPolicy, error) {
	var err error
	var rs []*app.AccessPolicy

	if rID == 0 {
		rs, err = repo.ListAll(ctx, ids)
	} else {
		rs, err = repo.ListByRoleID(ctx, rID)
	}

	return rs, err
}

// ListAll is to list access policies.
func (repo *AccessPolicy) ListAll(ctx context.Context, ids []uint64) ([]*app.AccessPolicy, error) {
	query, args, err := sql.In(AccessPolicyListQuery, ids)
	if err != nil {
		repo.logger.Log("method", "listAll", "error", err)
		return nil, err
	}
	query = sql.Rebind(DriverName, query)
	var aa []*accessPolicyAttrib

	{
		if err := repo.db.Select(&aa, query, args...); err != nil {
			repo.logger.Log("method", "listAll", "error", err)
			return nil, err
		}
	}
	var pp []*app.AccessPolicy
	{
		for _, a := range aa {
			pp = append(pp, a.toAccessPolicy())
		}
	}

	return pp, nil
}

// Get a AccessPolicy.
func (repo *AccessPolicy) Get(ctx context.Context, id uint64) (*app.AccessPolicy, error) {
	var t []accessPolicyAttrib

	if err := repo.db.Select(&t, AccessPolicyGetQuery, id); err != nil {
		repo.logger.Log("method", "get", "error", err)
		return nil, err
	}

	if len(t) == 0 {
		return nil, nil
	}

	return t[0].toAccessPolicy(), nil
}

// ListByRoleID gets a list of access_policies for a role.
func (repo *AccessPolicy) ListByRoleID(ctx context.Context, id uint64) ([]*app.AccessPolicy, error) {
	query, args, err := sql.In(AccessPolicyListByRoleIDQuery, id)
	if err != nil {
		repo.logger.Log("method", "listByRoleID", "error", err)
		return nil, err
	}

	var aa []*accessPolicyAttrib
	{
		if err := repo.db.Select(&aa, query, args...); err != nil {
			repo.logger.Log("method", "listByRoleID", "error", err)
			return nil, err
		}
	}

	var aps []*app.AccessPolicy
	{
		for _, v := range aa {
			aps = append(aps, v.toAccessPolicy())
		}
	}
	return aps, nil
}

// ListByUserPermission gets a list of roles for a user.
func (repo *AccessPolicy) ListByUserPermission(ctx context.Context, userID, permissionID uint64) ([]uint64, error) {
	query, args, err := sql.In(AccessPolicyListByUserPermissionQuery, permissionID, userID)
	if err != nil {
		repo.logger.Log("method", "listByUserPermission", "error", err)
		return nil, err
	}
	var as []uint64
	{
		if err := repo.db.Select(&as, query, args...); err != nil {
			repo.logger.Log("method", "listByUserPermission", "error", err)
			return nil, err
		}
	}
	return as, nil
}

// Delete is to delete access policies.
func (repo *AccessPolicy) Delete(ctx context.Context, id uint64) error {
	_, err := repo.db.Exec(AccessPolicyDeleteQuery, id)
	return err
}

// Update is to update a given access policy.
func (repo *AccessPolicy) Update(ctx context.Context, ap *app.AccessPolicy) (*app.AccessPolicy, error) {
	a := toAccessPolicyAttribs(ap)
	if err := repo.update(ctx, a); err != nil {
		repo.logger.Log("method", "update", "error", err)
		return nil, err
	}

	return a.toAccessPolicy(), nil
}

func (repo *AccessPolicy) update(ctx context.Context, a *accessPolicyAttrib, args ...string) error {
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
		s, err := sql.PartialUpdateStmt(a, "access_policies", condition, repo.fieldInfo)

		if err != nil {
			return err
		}

		setStatement = s
	}

	_, err := repo.db.Exec(setStatement)
	return err
}

func toAccessPolicyAttribs(ap *app.AccessPolicy) *accessPolicyAttrib {
	a := &accessPolicyAttrib{}

	if ap == nil {
		return a
	}

	a.ID = ap.Id
	a.Name = ap.Name
	a.State = ap.State

	return a
}

func (a *accessPolicyAttrib) toAccessPolicy() *app.AccessPolicy {
	ap := &app.AccessPolicy{}
	if a == nil {
		return ap
	}

	ap.Id = a.ID
	ap.Name = a.Name
	ap.State = a.State
	ap.CreatedAt = typeutils.ToUnixTime(a.CreatedAt)
	ap.UpdatedAt = typeutils.ToUnixTime(a.UpdatedAt)

	return ap
}

func (repo *AccessPolicy) validate(a *accessPolicyAttrib, validateOnlyPresent ...bool) error {
	_, err := repo.fieldInfo.Validate(a, validateOnlyPresent...)

	if err == nil {
		return err
	}

	return nil
}
