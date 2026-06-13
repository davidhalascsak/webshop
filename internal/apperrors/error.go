package apperrors

import (
	"errors"
)

var (
	ErrBadRequest   = errors.New("bad_request")
	ErrNotFound     = errors.New("not_found")
	ErrConflict     = errors.New("conflict")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInternal     = errors.New("internal")
)

type AppError struct {
	Kind    error
	Message string
}

func (e *AppError) Error() string { return e.Message }
