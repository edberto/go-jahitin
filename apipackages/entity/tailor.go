package entity

import (
	"database/sql"
	"time"
)

type (
	TailorEntity struct {
		ID        int
		Email     string
		UUID      string
		Phone     sql.NullString
		Address   sql.NullString
		Name      sql.NullString
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
