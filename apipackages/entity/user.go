package entity

import (
	"database/sql"
	"time"
)

type (
	UserEntity struct {
		ID        int
		Role      int
		Phone     string
		Email     string
		Name      string
		Password  string
		Username  string
		UUID      string
		TailorID  sql.NullInt32
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
