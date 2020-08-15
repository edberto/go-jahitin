package request

type (
	Register struct {
		Email    string `json:"email" validate:"required,email"`
		Phone    string `json:"phone" validate:"required,startswith=08,min=10,max=13"`
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required,rolecheck"`
	}
)
