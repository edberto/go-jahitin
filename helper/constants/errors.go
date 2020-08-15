package constants

import "errors"

var (
	ErrIncorrectPassword   = errors.New("Incorrect Password")
	ErrUsernameHasBeenUsed = errors.New("Username has been used")
)
