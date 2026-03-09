package storage

import "errors"

var (
	ErrUserExist    = errors.New("user already exists")
	ErrUserNotFound = errors.New("user noy found")
	ErrAppNotFound  = errors.New("app not found")
)
