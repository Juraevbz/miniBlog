package errs

import "errors"

var (
	ErrBadRequest = errors.New("ErrBadRequest")
	ErrIntervalServerError = errors.New("ErrIntervalServerError")
)