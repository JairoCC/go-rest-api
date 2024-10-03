package model

import "errors"

var (
	ErrPersonCannotBeNil    = errors.New("person cannot be null")
	ErrIDPersonDoesNotExist = errors.New("person does not exist")
)
