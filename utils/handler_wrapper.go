package utils

import (
	"encoding/json"
	"github.com/ansel1/merry"
	"github.com/gorilla/mux"
	"net/http"
)

// WrapHandler Extract attributes of merry error and write them to ResponseWriter
func WrapHandler(handler func(request *http.Request, vars map[string]string) ([]byte, int, error)) func(writer http.ResponseWriter, request *http.Request) {
	// We want to return a function with the signature for an http handler
	f := func(writer http.ResponseWriter, request *http.Request) {
		// We call the handler we are wrapping and capture its return values
		buf, statusCode, err := handler(request, mux.Vars(request))

		logger := Logger(request.Context())
		if err != nil {
			logger.Error(err)

			var message string
			rootError := err
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
			err = merry.WithHTTPCode(err, 500).WithMessage("Unable to write to response writer")
			logger.WithError(err).Error(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
	return f
}

func getHTTPStatus(err error) int {
	code := merry.HTTPCode(err)
	return code
}
