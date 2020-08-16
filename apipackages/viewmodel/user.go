package viewmodel

type (
	UserVM struct {
		ID       int `json:"id"`
		TailorID int `json:"tailor_id"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Username string `json:":"username"`
		UUID     string `json:"uuid"`
	}
)
