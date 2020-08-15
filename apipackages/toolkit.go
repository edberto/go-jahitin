package apipackages

import (
	"database/sql"
	"go-jahitin/helper/auth"

	"github.com/go-playground/validator/v10"
)

type (
	Toolkit struct {
		DB          *sql.DB
		Validator   *validator.Validate
		AccessAuth  auth.IAuth
		RefreshAuth auth.IAuth
	}
)
