package errors

import (
	"strconv"

	"github.com/ansel1/merry"
)

const (
	errorKey      = "error"
	messageKey    = "message"
	stackTraceKey = "stackTrace"
	maxDepth      = 5
)

// AsFields Turns an error into a list of loggable fields
func AsFields(err error) map[string]interface{} {
	return asFields(err, 0)
}

func asFields(err error, level int) map[string]interface{} {
	if err == nil {
		return make(map[string]interface{})
	}
	if level > maxDepth {
		return map[string]interface{}{"AsFieldsError": "Maximum recursion depth exceeded"}
	}

	fields := make(map[string]interface{})

	if mErr, ok := err.(merry.Error); ok {
		// set outermost stacktrace in the fields map first so it can be overwritten by a nested error if present
		fields[stackTraceKey] = merry.Stacktrace(err)
		for k, v := range merry.Values(mErr) {
			if key, ok := k.(string); ok {
				if mErr, ok := v.(merry.Error); ok {
					nestedFields := asFields(mErr, level+1)
					for nestedK, nestedV := range nestedFields {
						fields[nestedK] = nestedV
					}
				}
				// don't namespace stacktrace by depth, only need one
				if key != stackTraceKey && level > 0 {
					key = key + "_" + strconv.Itoa(level)
				}
				fields[key] = v
			}
		}
		fields[messageKey] = merry.Message(err)
	} else {
		fields[messageKey] = err.Error()
	}
	fields[errorKey] = err.Error()

	return fields
}
