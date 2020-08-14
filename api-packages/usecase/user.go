package usecase

import (
	"go-jahitin/api-packages/model"
	"go-jahitin/api-packages/viewmodel"

	"github.com/gin-gonic/gin"
)

type (
	IUser interface {
		Register(c *gin.Context, param RegisterUserParam) (viewmodel.UserVM, error)
		GetOne(c *gin.Context, param GetOneUserParam) (viewmodel.UserVM, error)
	}

	User struct {
		UserModel model.IUser
	}

	RegisterUserParam struct{}
	GetOneUserParam   struct{}
)

func NewUserUC() IUser {
	return &User{
		UserModel: model.NewUserModel(),
	}
}

func (uc *User) Register(c *gin.Context, param RegisterUserParam) (viewmodel.UserVM, error) {
	return viewmodel.UserVM{}, nil
}

func (uc *User) GetOne(c *gin.Context, param GetOneUserParam) (viewmodel.UserVM, error) {
	return viewmodel.UserVM{}, nil
}
