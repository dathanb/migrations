package utils

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/udacity/go-errors"
	"github.com/ansel1/merry"
	"encoding/json"
)

// WrapHandler Extract attributes of merry error and write them to ResponseWriter
func WrapHandler(handler func(request *http.Request, vars map[string]string) ([]byte, int, error)) func(writer http.ResponseWriter, request *http.Request) {
	// We want to return a function with the signature for an http handler
	f := func(writer http.ResponseWriter, request *http.Request) {
		// We call the handler we are wrapping and capture its return values
		buf, statusCode, err := handler(request, mux.Vars(request))

		logger := Logger(request.Context())
		if err != nil {
			logger.WithFields(errors.AsFields(err)).Error(err)

			var message string
			rootError := errors.RootCause(err)
			if rootError != nil {
				message = merry.Message(rootError)
			} else {
				message = merry.Message(err)
			}
			responseBody := MakeErrorResponse(message)

			statusCode = getHTTPStatus(err)
			buf, _ = json.Marshal(&responseBody)
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(statusCode)
		_, err = writer.Write(buf)

		if err != nil {
			err = errors.WithRootCause(errors.UnexpectedError, err).WithMessage("Unable to write to response writer")
			logger.WithFields(errors.AsFields(err)).WithError(err).Error(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
	return f
}

func getHTTPStatus(err error) int {
	code := merry.HTTPCode(err)
	if errors.IsCausedBy(err, errors.ArgumentError) {
		code = http.StatusBadRequest
	} else if errors.IsCausedBy(err, errors.SQLUniquenessConstraintError) {
		code = http.StatusConflict
	} else if errors.IsCausedBy(err, errors.SQLConstraintViolationError) {
		code = http.StatusConflict
	} else if errors.IsCausedBy(err, errors.RequestBodyError) {
		code = http.StatusBadRequest
	}
	return code
}
