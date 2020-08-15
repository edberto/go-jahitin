package entity

import "time"

type (
	UserTokenEntity struct {
		UserID int
		UUID string
		ExpiredAt time.Time
	}
)