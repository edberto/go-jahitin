package request

type (
	InsertBulkTailorMaterial struct {
		TailorID   int     `json:"tailor_id" validate:"required"`
		MaterialID int     `json:"material_id" validate:"required"`
		Price      float64 `json:"price"`
	}
)
