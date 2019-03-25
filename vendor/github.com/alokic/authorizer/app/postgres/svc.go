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
	// SvcListAllQuery is list query template for listing Svcs.
	SvcListAllQuery = "SELECT * FROM services"
	// SvcGetQuery is lightest query template to get a Svc.
	SvcGetQuery = "SELECT * FROM services WHERE id = $1"
)

// Svc is postgres repo for Svcs.
type Svc struct {
	db        DB
	logger    Logger
	fieldInfo *sql.FieldInfo
}

// svcAttrib desines attributes of
type svcAttrib struct {
	ID                 uint64    `db:"id"`
	Name               string    `db:"name"`
	Domain             string    `db:"domain"`
	State              string    `db:"state"`
	AuthenticationType string    `db:"authentication_type"`
	AuthenticationKey  string    `db:"authentication_key"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// NewSvc is Svc constructor.
func NewSvc(d DB, logger Logger) *Svc {
	return &Svc{db: d, logger: logger, fieldInfo: sql.GenFieldInfo(DriverName, svcAttrib{})}
}

// Create a Svc in db.
func (repo *Svc) Create(ctx context.Context, e *app.Svc) (*app.Svc, error) {
	var t *svcAttrib
	{
		t = toSvcAttribs(e)
		t.CreatedAt = time.Now().UTC()
		t.UpdatedAt = t.CreatedAt

		err := repo.validate(t)
		if err != nil {
			repo.logger.Log("method", "create", "error", err)
			return nil, err
		}
	}

	err := BatchInsert(repo.db, "svcs", []interface{}{t}, repo.fieldInfo)
	if err != nil {
		repo.logger.Log("method", "create", "error", err)
		return nil, err
	}

	return t.toSvc(), nil
}

// Get a Svc.
func (repo *Svc) Get(ctx context.Context, id uint64) (*app.Svc, error) {
	var t []svcAttrib

	if err := repo.db.Select(&t, SvcGetQuery, id); err != nil {
		repo.logger.Log("method", "get", "error", err)
		return nil, err
	}

	if len(t) == 0 {
		return nil, nil
	}

	return t[0].toSvc(), nil
}

// ListAll lists all svcs from the db.
func (repo *Svc) ListAll(ctx context.Context) ([]*app.Svc, error) {
	var pa []*svcAttrib
	{
		if err := repo.db.Select(&pa, SvcListAllQuery); err != nil {
			repo.logger.Log("method", "listall", "error", err)
			return nil, err
		}
	}
	var ps []*app.Svc
	{
		for _, a := range pa {
			ps = append(ps, a.toSvc())
		}
	}

	return ps, nil
}

// Update a svc in the db
func (repo *Svc) Update(ctx context.Context, e *app.Svc) (*app.Svc, error) {
	p := toSvcAttribs(e)

	if err := repo.update(ctx, p); err != nil {
		return nil, err
	}

	return p.toSvc(), nil
}

func (repo *Svc) update(ctx context.Context, p *svcAttrib, args ...string) error {
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
		s, err := sql.PartialUpdateStmt(p, "services", condition, repo.fieldInfo)
		if err != nil {
			return err
		}
		setStatement = s
	}

	_, err := repo.db.Exec(setStatement)
	return err
}

func toSvcAttribs(t *app.Svc) *svcAttrib {
	e := &svcAttrib{}
	if t == nil {
		return e
	}
	e.ID = t.Id
	e.Name = t.Name
	e.Domain = t.Domain
	e.State = t.State

	return e
}

func (repo *Svc) validate(p *svcAttrib, validateOnlyPresent ...bool) error {
	_, err := repo.fieldInfo.Validate(p, validateOnlyPresent...)

	if err == nil {
		return err
	}

	return nil
}

func (t *svcAttrib) toSvc() *app.Svc {
	e := &app.Svc{}
	if t == nil {
		return e
	}

	e.Id = t.ID
	e.Name = t.Name
	e.Domain = t.Domain
	e.State = t.State
	e.CreatedAt = typeutils.ToUnixTime(t.CreatedAt)
	e.UpdatedAt = typeutils.ToUnixTime(t.UpdatedAt)
	e.AuthenticationType = t.AuthenticationType
	e.AuthenticationKey = t.AuthenticationKey

	return e
}
