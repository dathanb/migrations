package error

type Error interface {
	error
}

type genericError struct {
	msg string
}

func (e *genericError) Error() string {
	return e.msg
}

func NotImplemented() Error {
	return &genericError{msg: "Not implemented"}
}

func WithMessage(msg string) Error {
	return &genericError{msg: msg}
}

