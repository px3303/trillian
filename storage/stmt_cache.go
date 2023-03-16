package storage

import (
	"context"
	"database/sql"
)

// Stmt is a prepared statement.
// A Stmt is safe for concurrent use by multiple goroutines.
type Stmt interface {
	// Close closes the statement.
	Close() error

	// ExecContext executes a prepared statement with the given arguments and
	// returns a Result summarizing the effect of the statement.
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)

	// QueryContext executes a prepared query statement with the given arguments
	// and returns the query results as a *Rows.
	QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)

	// QueryRowContext executes a prepared query statement with the given arguments.
	// If an error occurs during the execution of the statement, that error will
	// be returned by a call to Scan on the returned *Row, which is always non-nil.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.
	QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row

	// WithTxContext returns a transaction-specific prepared statement from
	// an existing statement.
	WithTxContext(ctx context.Context, tx *sql.Tx) Stmt
}

// StmtCache is a Stmt caches
type StmtCache interface {
	// getStmt creates and caches Stmt based on the passed in statement and number of bound arguments.
	GetStmt(ctx context.Context, statement string, num int, first, rest string) Stmt
}
