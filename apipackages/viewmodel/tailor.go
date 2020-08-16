package viewmodel

type (
	TailorVM struct {
		ID            int     `json:"id"`
		ModelID       int     `json:"model_id"`
		MaterialID    int     `json:"materila_id"`
		ModelName     string  `json:"model_name"`
		MaterialName  string  `json:"material_name"`
		MaterialColor string  `json:"material_color"`
		UUID          string  `json:"uuid"`
		Name          string  `json:"name"`
		Phone         string  `json:"phone"`
		Email         string  `json:"email"`
		Address       string  `json:"address"`
		Price         float64 `json:"price"`
	}
)
