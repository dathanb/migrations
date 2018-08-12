package errors

import "github.com/ansel1/merry"

var (
	// MarshalingError wraps errors that occur during marshaling
	MarshalingError = merry.New("Marshaling error")

	// JSONMarshalingError indicates an issue with masrhaling a struct to json
	JSONMarshalingError = merry.WithMessage(MarshalingError, "JSON marshaling error")
)
