package entity

import "time"

type (
	TailorMaterialEntity struct {
		ID         int
		TailorID   int
		MaterialID int
		Price      float64
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
