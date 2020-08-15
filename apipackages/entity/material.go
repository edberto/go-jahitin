package entity

import "time"

type (
	MaterialEntity struct {
		ID        int
		UUID      string
		Detail    string
		Name      string
		Color     string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
