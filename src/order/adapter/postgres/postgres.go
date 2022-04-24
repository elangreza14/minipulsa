package postgres

import (
	"context"
	"database/sql"
)

type (
	// in Hex ARCH
	// generating available method into DB postgres
	// co we can focus in adapting this postgres adaptor into weight repository
	DBTX interface {
		ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
		QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
		QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	}

	Queries struct {
		db DBTX
	}
)

// New creating Queries instance
func New(db DBTX) *Queries {
	return &Queries{
		db: db,
	}
}
