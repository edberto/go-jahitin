package entity

import "time"

type (
	ModelEntity struct {
		ID        int
		UUID      string
		Detail    string
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
