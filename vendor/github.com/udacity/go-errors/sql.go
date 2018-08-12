package errors

import (
	"net/http"

	"github.com/ansel1/merry"
)

var (
	// SQLError wraps errors from the db access library
	SQLError = merry.New("SQL error")

	// SQLCommitError occurs when a commit failed
	SQLCommitError = merry.WithMessage(SQLError, "Commit failure")

	// SQLPrepareError wraps errors in preparing a SQL statement
	SQLPrepareError = merry.WithMessage(SQLError, "SQL syntax error")

	// SQLSelectError indicates error selecting data after statement is prepared
	SQLSelectError = merry.WithMessage(SQLError, "SQL select error")

	// SQLInsertError indicates error inserting data after statement is prepared
	SQLInsertError = merry.WithMessage(SQLError, "SQL insert error")

	// SQLDeleteError indicates error deleting data after statement is prepared
	SQLDeleteError = merry.WithMessage(SQLError, "SQL delete error")

	// SQLUpdateError indicates error updated data after statement is prepared
	SQLUpdateError = merry.WithMessage(SQLError, "SQL update error")

	// SQLConstraintViolationError occurs when an insert or update violates a contraint (e.g. FK, uniqueness, non-null)
	SQLConstraintViolationError = merry.WithMessage(SQLError, "Constraint violation")

	// SQLUniquenessConstraintError occurs when an insert or update violates a uniqueness contraint
	SQLUniquenessConstraintError = merry.WithMessage(SQLConstraintViolationError, "SQL uniqueness constraint violated").WithHTTPCode(http.StatusConflict)
)
