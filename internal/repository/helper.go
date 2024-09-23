package repository

import (
	"errors"
	"mini_blog/internal/errs"

	"gorm.io/gorm"
)

func handleError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}
	return err
}
