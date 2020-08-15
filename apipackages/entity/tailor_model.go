package entity

import "time"

type (
	TailorModelEntity struct {
		ID        int
		TailorID  int
		ModelID   int
		Price     float64
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
