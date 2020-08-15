package viewmodel

type (
	OrderVM struct {
		ID            int           `json:"id"`
		UserID        int           `json:"user_id"`
		TailorID      int           `json:"tailor_id"`
		UserName      string        `json:"user_name"`
		UserPhone     string        `json:"user_phone"`
		UserAddress   string        `json:"user_address"`
		TailorName    string        `json:"tailor_name"`
		TailorPhone   string        `json:"tailor_phone"`
		Status        string        `json:"status"`
		UUID          string        `json:"uuid"`
		Specification Specification `json:"specification"`
	}

	Specification struct {
		MaterialID     int      `json:"material_id"`
		ModelID        int      `json:"model_id"`
		MaterialName   string   `json:"material_name"`
		MaterialColor  string   `json:"material_color"`
		MaterialDetail string   `json:"material_detail"`
		ModelName      string   `json:"model_name"`
		ModelDetail    string   `json:"model_detail"`
		Quantity       Quantity `json:"quantity"`
	}

	Quantity struct {
		XS  int `json:"xs"`
		S   int `json:"s"`
		M   int `json:"m"`
		L   int `json:"l"`
		XL  int `json:"xl"`
		XXL int `json:"xxl"`
		LLL int `json:"lll"`
	}
)
