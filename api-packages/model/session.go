package model

import (
	"go-jahitin/api-packages/entity"

	"github.com/gin-gonic/gin"
)

type (
	ISession interface {
		InsertOne(c *gin.Context, param InsertOneSessionParam) (entity.SessionEntity, error)
		DeleteOne(c *gin.Context, param DeleteOneSessionParam) error
	}

	Session struct{}

	InsertOneSessionParam struct{}
	DeleteOneSessionParam struct{}
)

func NewSessionModel() ISession {
	return &Session{}
}

func (model *Session) InsertOne(c *gin.Context, param InsertOneSessionParam) (entity.SessionEntity, error) {
	return entity.SessionEntity{}, nil
}

func (model *Session) DeleteOne(c *gin.Context, param DeleteOneSessionParam) error {
	return nil
}
