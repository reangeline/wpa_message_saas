package erro

import "errors"

var (
	ErrValidEmail         = errors.New("please add a valid email")
	ErrEmailIsRequired    = errors.New("email is required")
	ErrNameIsRequired     = errors.New("name is required")
	ErrLastNameIsRequired = errors.New("lastname is required")
)
