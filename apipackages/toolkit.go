package apipackages

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type (
	Toolkit struct {
		DB        *sql.DB
		Validator *validator.Validate
	}
)
