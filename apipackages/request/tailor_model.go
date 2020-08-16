package request

type (
	InsertBulkTailorModel struct {
		TailorID int     `json:"tailor_id" validate:"required"`
		ModelID  int     `json:"model_id" validate:"required"`
		Price    float64 `json:"price"`
	}
)
