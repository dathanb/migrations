package errors

import (
	"net/http"

	"github.com/ansel1/merry"
)

var (
	// InputError wraps data input errors
	InputError = merry.New("Input error")

	// ArgumentError error with one or more input arguments
	ArgumentError = merry.WithMessage(InputError, "Invalid argument error").WithHTTPCode(http.StatusBadRequest)

	// RequestBodyError error with the request body
	RequestBodyError = merry.WithMessage(InputError, "Invalid request body errors").WithHTTPCode(http.StatusBadRequest)


	// StatusError error with conflicting statuses
	StatusConflictError = merry.WithMessage(InputError, "Conflicting status error").WithHTTPCode(http.StatusConflict);
)
