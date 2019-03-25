package postgres

import (
	"context"
	"time"

	"github.com/alokic/authorizer/app"
	sql "github.com/alokic/gopkg/sql/v2"
	"github.com/alokic/gopkg/typeutils"
)

var (
	// AccessPlcyPermDeleteQuery is used to delete association with matching access policy id
	AccessPlcyPermDeleteQuery = "DELETE FROM access_policies_permissions where access_policy_id = $1"
)

// AccessPlcyPermRepo is the postgres repo for AccessPolicyPermissions
type AccessPlcyPermRepo struct {
	db        DB
	logger    Logger
	fieldInfo *sql.FieldInfo
}

type accessPlcyPermAttribute struct {
	AccessPolicyID uint64    `db:"access_policy_id"`
	PermissionID   uint64    `db:"permission_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// NewAccessPlcyPerm is the constructor for the repo
func NewAccessPlcyPerm(d DB, logger Logger) *AccessPlcyPermRepo {
	return &AccessPlcyPermRepo{
		db:        d,
		logger:    logger,
		fieldInfo: sql.GenFieldInfo(DriverName, accessPlcyPermAttribute{}),
	}
}

// BatchUpdate deleted existing associations for the access policy and adds new associations for accesspolicy permissions
func (repo *AccessPlcyPermRepo) BatchUpdate(ctx context.Context, ap []*app.AccessPlcyPerm, apID uint64) ([]*app.AccessPlcyPerm, error) {
	var aa []*accessPlcyPermAttribute
	for _, pp := range ap {
		a := toAccessPolicyPermAttribs(pp)
		err := repo.validate(a)
		if err != nil {
			repo.logger.Log("method", "batchUpdate", "error", err)
			return nil, err
		}

		aa = append(aa, a)
	}

	repo.db.Transaction(func(tx *sql.Tx) (interface{}, error) {
		_, err := tx.Exec(AccessPlcyPermDeleteQuery, apID)
		if err != nil {
			return nil, err
		}

		var ii []interface{}
		for _, i := range aa {
			ii = append(ii, i)
		}

		stmt, params := sql.BatchInsertStatement("access_policies_permissions", ii, repo.fieldInfo)
		return tx.Exec(stmt, params...)
	})

	var apps []*app.AccessPlcyPerm
	for _, v := range aa {
		apps = append(apps, v.toAccessPolicy())
	}
	return apps, nil
}

func toAccessPolicyPermAttribs(a *app.AccessPlcyPerm) *accessPlcyPermAttribute {
	p := &accessPlcyPermAttribute{}
	if a == nil {
		return p
	}
	p.AccessPolicyID = a.AccessPolicyId
	p.PermissionID = a.PermissionId
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = p.CreatedAt

	return p
}

func (a *accessPlcyPermAttribute) toAccessPolicy() *app.AccessPlcyPerm {
	p := &app.AccessPlcyPerm{}
	if a == nil {
		return p
	}

	p.PermissionId = a.PermissionID
	p.AccessPolicyId = a.AccessPolicyID
	p.CreatedAt = typeutils.ToUnixTime(a.CreatedAt)
	p.UpdatedAt = typeutils.ToUnixTime(a.UpdatedAt)

	return p
}

func (repo *AccessPlcyPermRepo) validate(p *accessPlcyPermAttribute, validateOnlyPresent ...bool) error {
	_, err := repo.fieldInfo.Validate(p, validateOnlyPresent...)

	if err == nil {
		return err
	}

	return nil
}
