package model

import (
	"go-jahitin/api-packages/entity"

	"github.com/gin-gonic/gin"
)

type (
	IUser interface {
		InsertOne(c *gin.Context, param InsertOneUserParam) (entity.UserEntity, error)
		GetOne(c *gin.Context, param GetOneUserParam) (entity.UserEntity, error)
	}

	User struct{}

	InsertOneUserParam struct{}
	GetOneUserParam    struct{}
)

func NewUserModel() IUser {
	return &User{}
}

func (model *User) InsertOne(c *gin.Context, param InsertOneUserParam) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}

func (model *User) GetOne(c *gin.Context, param GetOneUserParam) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}
