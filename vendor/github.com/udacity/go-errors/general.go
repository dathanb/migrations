package errors

import (
	"github.com/ansel1/merry"
)

const (
	rootCauseKey = "RootCause"
)

var (
	// UnexpectedError wraps general, unplanned or unhandled errors
	UnexpectedError = merry.New("Unexpected error")

	// UnexpectedConfigurationError wraps errors with configuration
	UnexpectedConfigurationError = merry.WithMessage(UnexpectedError, "Unexpected configuration error")

	// UnexpectedDatabaseError wraps unexpected errors that hapenned in the data access layer
	UnexpectedDatabaseError = merry.WithMessage(UnexpectedError, "Unexpected data access or database error")

	// NotImplementedError reports a method not yet implemented
	NotImplementedError = merry.New("Method not implemented")
)

// WithRootCause captures the stack trace one step above this function and adds error as RootCause value
func WithRootCause(merryError merry.Error, rootCause error) merry.Error {
	return merryError.WithStackSkipping(1).WithValue(rootCauseKey, rootCause)
}

// RootCause extracts root cause error from merry Error values
func RootCause(err error) error {
	cause := merry.Value(err, rootCauseKey)
	if cause == nil {
		return nil
	}
	if mErr, success := cause.(merry.Error); success {
		result := RootCause(mErr)
		if result != nil {
			cause = result
		}
	}
	return cause.(error)
}

// IsCausedBy checks if an error was caused by the specified error at any point
func IsCausedBy(err error, originals ...error) bool {
	if merry.Is(err, originals...) {
		return true
	}

	cause := merry.Value(err, rootCauseKey)

	if cause == nil {
		return false
	}
	return IsCausedBy(cause.(error), originals...)
}
