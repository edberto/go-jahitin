package handler

import (
	"go-jahitin/api-packages/usecase"

	"github.com/gin-gonic/gin"
)

type (
	IUser interface {
		Register(c *gin.Context)
		GetOne(c *gin.Context)
	}

	User struct {
		UserUC usecase.IUser
	}
)

func NewUserHandler() IUser {
	return &User{
		UserUC: usecase.NewUserUC(),
	}
}

func (h *User) Register(c *gin.Context) {
	return
}

func (h *User) GetOne(c *gin.Context) {
	return
}
