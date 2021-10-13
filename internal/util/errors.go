package util

import "errors"

var (
	ErrUnknownArgument = errors.New("unknown argument passed")

	ErrInvalidArgument = errors.New("invalid argument passed")
)
