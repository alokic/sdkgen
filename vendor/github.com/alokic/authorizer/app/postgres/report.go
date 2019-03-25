package postgres

import (
	"context"

	sql "github.com/alokic/gopkg/sql/v2"

	"github.com/alokic/authorizer/app"
)

// Report is postgres repo for reports.
type Report struct {
	db     *sql.DB
	logger Logger
}

// NewReport is constructor for reports.
func NewReport(d *sql.DB, logger Logger) *Report {
	return &Report{db: d, logger: logger}
}

// Ping checks whether database(s) are reaxchable.
func (r *Report) Ping(ctx context.Context, args ...interface{}) (*app.Report, error) {
	err := r.db.Ping()
	if err != nil {
		return &app.Report{Svc: "postgres", Err: err.Error()}, err
	}
	return &app.Report{Svc: "postgres", Msg: "OK"}, nil
}
