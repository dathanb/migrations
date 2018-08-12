package errors

import (
	"net/http"

	"github.com/ansel1/merry"
)

var (
	// HTTPError is the base error for any errors related to http client calls
	HTTPError = merry.New("HTTP error")

	// HTTPBadRequestError is for 400 error codes
	HTTPBadRequestError = merry.WithMessage(HTTPError, "400 Bad Request status code received").WithHTTPCode(http.StatusBadRequest)

	// HTTPUnauthorizedError is for 401 error codes
	HTTPUnauthorizedError = merry.WithMessage(HTTPBadRequestError, "401 Unauthorized HTTP status code received").WithHTTPCode(http.StatusUnauthorized)

	// HTTPForbiddenError is for 403 error codes
	HTTPForbiddenError = merry.WithMessage(HTTPBadRequestError, "403 Forbidden HTTP status code received").WithHTTPCode(http.StatusForbidden)

	// HTTPNotFoundError for 404 error codes
	HTTPNotFoundError = HTTPError.WithMessage("404 Not Found HTTP status code received").WithHTTPCode(http.StatusNotFound)

	// HTTPUnprocessableError if for 422 error codes
	HTTPUnprocessableError = merry.WithMessage(HTTPBadRequestError, "422 Unprocessable Entity HTTP status code received").WithHTTPCode(http.StatusUnprocessableEntity)
)
