package errors

const (
	// PqUniqueViolation is the name for a postresql uniqueness constraint error
	PqUniqueViolation string = "unique_violation"
)

const (
	// PqConstraintViolation is the class of all errors of constraint violations
	PqConstraintViolation string = "23"
)
