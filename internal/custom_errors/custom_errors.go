package custom_errors

import "errors"

var (
	ErrCredentialsInUse = errors.New("username already in use")
	ErrWrongCredentials = errors.New("no such user or wrong password")
	ErrNotFound         = errors.New("not found")
)
