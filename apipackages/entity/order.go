package entity

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type (
	OrderEntity struct {
		ID            int
		UserID        int
		TailorID      int
		Status        int
		Price         float64
		UUID          string
		CreatedAt     time.Time
		UpdatedAt     time.Time
		UserAddress   sql.NullString
		Specification Specification
	}

	Specification struct {
		MaterialID int `json:"material_id"`
		ModelID    int `json:"model_id"`
		Qty        Qty `json:"qty"`
	}

	Qty struct {
		XS  int `json:"xs"`
		S   int `json:"s"`
		M   int `json:"m"`
		L   int `json:"l"`
		XL  int `json:"xl"`
		XXL int `json:"xxl"`
		LLL int `json:"lll"`
	}
)

func (s Specification) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Specification) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &s)
}

func (s Qty) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Qty) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &s)
}
