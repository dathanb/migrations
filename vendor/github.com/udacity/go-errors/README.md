# go-errors
A wrapper on Merry errors which includes deep IsCausedBy and WithCause features.

## Usage

### Logging a specific error with a root cause

```go
if err != nil {
  return errors.WithRootCause(errors.UnexpectedDatabaseError, err)
}
```

### AsFields
Gets all merry values, stack traces, etc as a map[string]interface{} for logging with Logrus.  For Non-Merry errors this will return a map with just a message and error set.

```go
err = errors.UnexpectedError.Here()
log.WithFields(errors.AsFields(err)).Error(err)
// {error: ..., stackTrace: ..., message: ..., etc}
```
