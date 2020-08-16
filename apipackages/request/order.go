package request

type (
	UpdateStatusOrder struct {
		ID     int    `json:"id" validate:"required,numeric"`
		Status string `json:"status" validate:"required,statuscheck"`
	}

	InsertOrder struct {
		UserID     int     `json:"user_id" validate:"required,numeric"`
		TailorID   int     `json:"tailor_id" validate:"required,numeric"`
		MaterialID int     `json:"material_id" validate:"required,numeric"`
		ModelID    int     `json:"model_id" validate:"required,numeric"`
		XSQty      int     `json:"xs_qty"`
		SQty       int     `json:"s_qty"`
		MQty       int     `json:"m_qty"`
		LQty       int     `json:"l_qty"`
		XLQty      int     `json:"xl_qty"`
		XXLQty     int     `json:"xxl_qty"`
		LLLQty     int     `json:"lll_qty"`
		Price      float64 `json:"price" validate:"required,numeric"`
	}
)
