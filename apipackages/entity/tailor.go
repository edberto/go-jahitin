package entity

import "time"

type (
	TailorEntity struct {
		ID        int
		Address   string
		Email     string
		Name      string
		UUID      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
