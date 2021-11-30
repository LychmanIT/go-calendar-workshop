package util

import "errors"

var (
	ErrUnknownArgument = errors.New("unknown argument passed")

	ErrInvalidArgument = errors.New("invalid argument passed")

	ErrEventNotFound = errors.New("event not found")

	ErrUserNotFound = errors.New("user not found")

	ErrAlreadyLoggedIn = errors.New("you have already logged in")

	ErrUnauthorized = errors.New("unauthorized")
)
