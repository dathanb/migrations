package utils

// Error is the response body when errors are encountered
type Error struct {
	Message string `json:"message"`
}

// ErrorResponse is the response body when errors encountered
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

func MakeErrorResponse(message string) *ErrorResponse {
	err := Error{
		Message: message,
	}
	return &ErrorResponse{Errors: []Error{err}}
}

