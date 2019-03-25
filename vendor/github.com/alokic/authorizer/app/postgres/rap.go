package postgres

import (
	"context"
	"time"

	"github.com/alokic/authorizer/app"
	sql "github.com/alokic/gopkg/sql/v2"
	"github.com/alokic/gopkg/typeutils"
)

var (
	// RoleAccessPolicyDeleteQuery is the query to delete associations with matching access policies.
	RoleAccessPolicyDeleteQuery = "DELETE FROM role_access_policies WHERE access_policy_id = $1"
)

// RolePlcyRepo is the postgres for role_aceess_policies.
type RolePlcyRepo struct {
	db        DB
	logger    Logger
	fieldInfo *sql.FieldInfo
}

type rolePlcyAttrib struct {
	RoleID         uint64    `db:"role_id"`
	AccessPolicyID uint64    `db:"access_policy_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// NewRolePlcy is the constructor for the postgres repo.
func NewRolePlcy(d DB, logger Logger) *RolePlcyRepo {
	return &RolePlcyRepo{db: d, logger: logger, fieldInfo: sql.GenFieldInfo(DriverName, rolePlcyAttrib{})}
}

// BatchUpdate deletes existing associations for access policy and updates the new associations.
func (repo *RolePlcyRepo) BatchUpdate(ctx context.Context, raps []*app.RoleAccessPolicy, aID uint64) ([]*app.RoleAccessPolicy, error) {
	var rr []*rolePlcyAttrib
	{
		for _, rp := range raps {
			a := toRolePlcyAttribs(rp)
			a.CreatedAt = time.Now().UTC()
			a.UpdatedAt = a.CreatedAt

			err := repo.validate(a)
			if err != nil {
				repo.logger.Log("method", "batch_update", "error", err)
				return nil, err
			}

			rr = append(rr, a)
		}
	}

	repo.db.Transaction(func(tx *sql.Tx) (interface{}, error) {
		_, err := tx.Exec(RoleAccessPolicyDeleteQuery, aID)
		if err != nil {
			return nil, err
		}

		var r []interface{}
		for _, ir := range rr {
			r = append(r, ir)
		}

		stmt, params := sql.BatchInsertStatement("role_access_policies", r, repo.fieldInfo)
		return tx.Exec(stmt, params...)
	})

	var rap []*app.RoleAccessPolicy
	for _, rp := range rr {
		rap = append(rap, rp.toRolePlcy())
	}

	return rap, nil
}

func toRolePlcyAttribs(r *app.RoleAccessPolicy) *rolePlcyAttrib {
	rp := &rolePlcyAttrib{}
	if r == nil {
		return rp
	}

	rp.AccessPolicyID = r.AccessPolicyId
	rp.RoleID = r.RoleId
	return rp
}

func (r *rolePlcyAttrib) toRolePlcy() *app.RoleAccessPolicy {
	rp := &app.RoleAccessPolicy{}
	rp.RoleId = r.RoleID
	rp.AccessPolicyId = r.AccessPolicyID
	rp.CreatedAt = typeutils.ToUnixTime(r.CreatedAt)
	rp.UpdatedAt = typeutils.ToUnixTime(r.UpdatedAt)
	return rp
}

func (repo *RolePlcyRepo) validate(r *rolePlcyAttrib, validateOnlyPresent ...bool) error {
	_, err := repo.fieldInfo.Validate(r, validateOnlyPresent...)
	return err
}
