package domain

import "errors"

var (
	ErrInvalidObjectID = errors.New("invalid object id")
	ErrInternalServer  = errors.New("internal server error")
	ErrDataNotFound    = errors.New("data not found")
)
