package errs

import "errors"

var (
	ErrBadRequest            = errors.New("ErrBadRequest")
	ErrRecordNotFound        = errors.New("ErrRecordNotFound")
	ErrInternalDatabaseError = errors.New("ErrInternalDatabaseError")
	ErrValidationFailed      = errors.New("ErrValidationFailed")
	ErrUnauthorized          = errors.New("ErrUnauthorizd")
	ErrUserNotFound          = errors.New("ErrUserNotFound")
	ErrUserInvalidPassword   = errors.New("ErrUserInvalidPassword")
)
